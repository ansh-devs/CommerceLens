package natsutil

import (
	"bytes"
	"encoding/gob"
	"sync"

	"github.com/ansh-devs/commercelens/order-service/dto"
	"github.com/nats-io/nats.go"
)

type NATSComponent struct {
	// nmu is the lock from the component.
	nmu sync.Mutex
	// nc is the connection to NATS Streaming.
	nc *nats.Conn
	// name is the name of component.
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
func (n *NATSComponent) UserIdEncoder(userID string) (bytes.Buffer, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(userID); err != nil {
		return bytes.Buffer{}, err
	} else {
		return buf, nil
	}
}

func (n *NATSComponent) DecryptMsgToOrder(data []byte) (dto.NatsPurchaseOrder, error) {
	var model dto.NatsPurchaseOrder
	enc := gob.NewDecoder(bytes.NewReader(data))
	if err := enc.Decode(&model); err != nil {
		return dto.NatsPurchaseOrder{}, err
	} else {
		return model, nil
	}
}

func (n *NATSComponent) DecryptMsgToUser(data []byte) (dto.NatsUser, error) {
	var model dto.NatsUser
	enc := gob.NewDecoder(bytes.NewReader(data))
	if err := enc.Decode(&model); err != nil {
		return dto.NatsUser{}, err
	} else {
		return model, nil
	}
}

func (n *NATSComponent) SendOrderNotPlacedMail() {
	// not implemented yet...
}
