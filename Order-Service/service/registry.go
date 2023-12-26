package service

import (
	"fmt"
	"github.com/go-kit/log/level"
	"github.com/hashicorp/consul/api"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func (s *OrderService) RegisterService(addr *string) {
	ttl := time.Second * 1
	fmt.Println(ttl.String())
	client := &api.AgentServiceCheck{
		CheckID:                        "status_alive",
		Name:                           "ORDER-SERVICE",
		TTL:                            ttl.String(),
		TLSSkipVerify:                  true,
		DeregisterCriticalServiceAfter: ttl.String(),
	}

	port, err := strconv.Atoi(strings.Trim(*addr, ":"))
	if err != nil {
		_ = level.Error(s.logger).Log("err", err)
	}
	regis := &api.AgentServiceRegistration{
		ID:      "instance_" + strconv.Itoa(rand.Int()),
		Name:    "ORDER-SERVICE",
		Tags:    []string{"order"},
		Port:    port,
		Address: s.getLocalIP().String(),
		Meta:    map[string]string{"registered_at": time.Now().String()},
		Check:   client,
	}
	//if err := s; err != nil {
	if err := s.consulClient.Agent().ServiceRegister(regis); err != nil {
		_ = level.Error(s.logger).Log("err", err)
	}
}

func (s *OrderService) UpdateHealthStatus() {
	ticker := time.NewTicker(time.Millisecond * 850)
	for {
		if err := s.consulClient.Agent().UpdateTTL(
			"status_alive",
			"[MESSAGE]: working as intended",
			api.HealthPassing,
		); err != nil {
			_ = level.Error(s.logger).Log("err", err)
		}
		<-ticker.C
	}
}
