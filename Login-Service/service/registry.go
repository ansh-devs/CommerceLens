package service

import (
	"github.com/go-kit/log/level"
	"github.com/hashicorp/consul/api"
	"strconv"
	"strings"
	"time"
)

func (s *LoginService) RegisterService(addr *string) {
	ttl := time.Second * 2
	checkClient := &api.AgentServiceCheck{
		CheckID:                        "service_alive" + s.SrvID,
		Name:                           "LOGIN-SERVICE",
		TTL:                            ttl.String(),
		TLSSkipVerify:                  true,
		DeregisterCriticalServiceAfter: ttl.String(),
		Notes:                          "Agent alive",
	}

	port, err := strconv.Atoi(strings.Trim(*addr, ":"))
	if err != nil {
		_ = level.Error(s.logger).Log("err", err)
	}
	srvRegister := &api.AgentServiceRegistration{
		ID:      s.SrvID,
		Name:    "LOGIN-SERVICE",
		Tags:    []string{"login", "auth"},
		Port:    port,
		Address: s.getLocalIP().String(),
		Meta:    map[string]string{"registered_at": time.Now().String()},
		Check:   checkClient,
	}

	if err := s.ConsulClient.Agent().ServiceRegister(srvRegister); err != nil {
		_ = level.Error(s.logger).Log("err", err)
	}
}

func (s *LoginService) UpdateHealthStatus() {
	ticker := time.NewTicker(time.Second * 1)
	for {

		if err := s.ConsulClient.Agent().UpdateTTL(
			"service_alive"+s.SrvID,
			"[MESSAGE]: Agent reachable",
			api.HealthPassing,
		); err != nil {
			_ = level.Error(s.logger).Log("err", err)
		}
		<-ticker.C
	}
}
