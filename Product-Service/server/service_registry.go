package server

import (
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

func (srvc *ProductService) registerService(addr string) {
	client := &api.AgentServiceCheck{
		CheckID:                        "status_alive",
		TTL:                            "5s",
		Name:                           "PRODUCT_SERVICE",
		TLSSkipVerify:                  true,
		DeregisterCriticalServiceAfter: "2s",
	}

	port, err := strconv.Atoi(strings.Trim(addr, ":"))
	if err != nil {
		fmt.Println("port conversion failed")
	}
	regis := &api.AgentServiceRegistration{
		ID:      "_instance_" + strconv.Itoa(rand.Int()),
		Name:    "PRODUCT_SERVICE",
		Tags:    []string{"products"},
		Port:    port,
		Address: getLocalIP().String(),
		Meta:    map[string]string{"registered_at": time.Now().String()},
		Check:   client,
	}
	if err := srvc.consulClient.Agent().ServiceRegister(regis); err != nil {
		fmt.Println(err)
	}

}

func (srvc *ProductService) updateHealthStatus() {
	ticker := time.NewTicker(time.Second * 3)
	for {
		if err := srvc.consulClient.Agent().UpdateTTL("status_alive", "[MESSAGE]: working as intended", api.HealthPassing); err != nil {
			log.Fatal(err)
		}
		<-ticker.C
	}
}

func getLocalIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddress := conn.LocalAddr().(*net.UDPAddr)

	return localAddress.IP
}

func (srvc *ProductService) getloginservice() (addr string, port int) {
	services, err := srvc.consulClient.Agent().Services()
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range services {
		if v.Service == "LOGIN_SERVICE" {
			addr = v.Address
			port = v.Port
		}
	}
	return addr, port
}
