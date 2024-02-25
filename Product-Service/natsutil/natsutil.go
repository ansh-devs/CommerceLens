package natsutil

import (
	"bytes"
	"encoding/gob"
	"github.com/nats-io/nats.go"
	"sync"
)

type NATSComponent struct {
	// nmu is the lock from the nats component.
	nmu sync.Mutex
	// nc is the connection established to NATS Streaming.
	nc *nats.Conn
	// name is the gives of the nats component.
	name string
}

// NewNatsComponent returns the instance for the Component.
func NewNatsComponent(compName string) *NATSComponent {
	return &NATSComponent{name: compName}
}

// ConnectToNATS connects to the NATS server.
func (n *NATSComponent) ConnectToNATS(url string, options ...nats.Option) error {
	n.nmu.Lock()
	nc, err := nats.Connect(url, options...)
	if err != nil {
		return err
	}
	n.nc = nc
	defer n.nmu.Unlock()
	return err
}

// NATS returns the current NATS connection.
func (n *NATSComponent) NATS() *nats.Conn {
	n.nmu.Lock()
	defer n.nmu.Unlock()
	return n.nc
}

// GracefulShutdown closes the connection to the NATS server
func (n *NATSComponent) GracefulShutdown() error {
	n.NATS().Close()
	return nil
}

func (n *NATSComponent) Publish(subject string, payload interface{}) error {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(payload); err != nil {
		return err
	}
	if err := n.NATS().Publish(subject, buf.Bytes()); err != nil {
		return err
	} else {
		return nil
	}
}
