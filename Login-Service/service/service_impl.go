package service

import (
	"context"
	"fmt"
	"github.com/ansh-devs/ecomm-poc/login-service/dto"
	"github.com/ansh-devs/ecomm-poc/login-service/natsutil"
	"github.com/ansh-devs/ecomm-poc/login-service/repo"
	"github.com/ansh-devs/ecomm-poc/login-service/token"
	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
	"github.com/hashicorp/consul/api"
	"github.com/nats-io/nats.go"
	"github.com/opentracing/opentracing-go"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"net"
	"strconv"
)

type LoginService struct {
	repository   repo.Repository
	logger       log.Logger
	ConsulClient *api.Client
	trace        opentracing.Tracer
	SrvID        string
	nats         *natsutil.NATSComponent
}

// NewService constructor for the OrderService...
func NewService(rep repo.Repository, logger log.Logger, tracer opentracing.Tracer) *LoginService {
	client, err := api.NewClient(&api.Config{
		Address: "service-discovery:8500",
	})
	if err != nil {
		fmt.Println(err)
	}
	srvID := "instance_" + strconv.Itoa(rand.Intn(99))
	nc := natsutil.NewNatsComponent(srvID)
	err = nc.ConnectToNATS("nats://nats-srvr:4222", nil)
	return &LoginService{
		repository:   rep,
		logger:       log.With(logger, "layer", "service"),
		ConsulClient: client,
		nats:         nc,
		trace:        tracer,
		SrvID:        "instance_" + strconv.Itoa(rand.Intn(99)),
	}
}

func (s *LoginService) GetUserWithNats(ctx context.Context) {
	nc := s.nats.NATS()
	for {
		_, err := nc.Subscribe("user.getdetails", func(msg *nats.Msg) {
			str, _ := s.nats.DecryptMsgToUserId(msg.Data)
			user, _ := s.repository.GetUser(ctx, str)
			encodedUser, _ := s.nats.EncodeUser(user)
			err := msg.Respond(encodedUser.Bytes())
			if err != nil {
				fmt.Println(err)
			}
		})
		if err != nil {
			fmt.Println(err)
		}
	}
}
func (s *LoginService) RegisterUser(ctx context.Context, payload dto.RegisterUserRequest) (dto.RegisterUserResponse, error) {
	resp, err := s.repository.CreateUser(ctx, payload)
	if err != nil {
		return dto.RegisterUserResponse{
			Message: err.Error(),
			Status:  "failed",
		}, nil
	}
	return dto.RegisterUserResponse{
		Message: "user created successfully",
		User:    resp,
		Status:  "success",
	}, nil
}

func (s *LoginService) LoginUser(ctx context.Context, email, password string) (dto.LoginUserResponse, error) {
	model, err := s.repository.CheckUserByEmail(ctx, email)
	if err != nil {
		return dto.LoginUserResponse{
			Message: err.Error(),
			Status:  "failed",
		}, nil
	}
	notValid := bcrypt.CompareHashAndPassword([]byte(model.Password), []byte(password))
	if err != nil {
		return dto.LoginUserResponse{
			Message: "can't verify your password",
			Status:  "failed",
		}, nil
	}
	if notValid == nil {
		gentoken, err := token.GenerateToken(model.ID)
		if err != nil {
			return dto.LoginUserResponse{
				Message: fmt.Sprintf("can't create auth token %s", err.Error()),
				Status:  "failed",
			}, nil
		}
		return dto.LoginUserResponse{
			Message: "ok",
			Status:  "success",
			Token:   gentoken,
		}, nil
	} else {
		return dto.LoginUserResponse{
			Message: "entered password is incorrect",
			Status:  "failed",
		}, nil
	}
}
func (s *LoginService) GetUserDetails(ctx context.Context, accessToken string) (dto.GetUserDetailsResponse, error) {
	tokenDecrypter, err := token.TokenDecrypter(accessToken)
	if err != nil {
		return dto.GetUserDetailsResponse{
			Message: fmt.Sprintf("can't verify your identity : %s", err.Error()),
			Status:  "failed",
		}, nil
	}
	user, err := s.repository.GetUser(ctx, tokenDecrypter.UserId)
	if err != nil {
		return dto.GetUserDetailsResponse{
			Message: fmt.Sprintf("can't fetch your details : %s", err.Error()),
			Status:  "failed",
		}, nil
	}
	return dto.GetUserDetailsResponse{
		Message: "ok",
		Status:  "success",
		User:    user,
	}, nil
}

// getLocalIP - place dto wrapper function around the method that makes calls to the repo...
func (s *LoginService) getLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		_ = level.Error(s.logger).Log("err", "can't get local ip")
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}
