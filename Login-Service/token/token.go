package token

import (
	"fmt"
	"time"
)

const (
	KEY = "YELLOW SUBMARINE, BLACK WIZARDRY"
)

type Maker interface {
	CreateToken(userId string, duration time.Duration) (token string, error error)
	VerifyToken(token string) (*Payload, error)
}

func GenerateToken(userId string) (token string, error error) {
	tokengen, err := newTokenMaker(KEY)
	if err != nil {
		return "", err
	}
	return tokengen.createTokenGen(userId, time.Hour*24*7)
}

func TokenDecrypter(token string) (*Payload, error) {
	maker, err := newTokenMaker(KEY)
	if err != nil {
		return nil, fmt.Errorf("[ERROR]: failed to create maker interface for decrypt token %s", err)
	}
	return maker.decryptTokenGen(token)
}
