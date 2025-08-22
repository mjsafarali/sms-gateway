package services

import (
	"errors"
	"github.com/nats-io/nats.go"
)

type Nats struct {
	jsc nats.JetStreamContext
}

func NewNatsService(jsc nats.JetStreamContext) *Nats {
	return &Nats{
		jsc: jsc,
	}
}

func (n *Nats) AddStream(name string, subjects []string) error {
	_, err := n.jsc.AddStream(&nats.StreamConfig{
		Name:     name,
		Subjects: subjects,
	})
	if err != nil && !errors.Is(err, nats.ErrStreamNameAlreadyInUse) {
		return err
	}

	return nil
}

func (n *Nats) Publish(subject string, data []byte) error {
	_, err := n.jsc.Publish(subject, data)
	if err != nil {
		return err
	}

	return nil
}
