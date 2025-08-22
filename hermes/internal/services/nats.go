package services

import (
	"context"
	"encoding/json"

	"github.com/nats-io/nats.go"
)

var PublisherSrv Publisher

type NatsPublisher struct {
	js nats.JetStreamContext
}

func NewNatsPublisher(js nats.JetStreamContext) *NatsPublisher {
	return &NatsPublisher{
		js: js,
	}
}

func (n *NatsPublisher) Publish(ctx context.Context, msg Msg) error {
	if _, err := n.js.AddStream(
		&nats.StreamConfig{
			Name:     "WALLET",
			Subjects: []string{"wallet.*"},
		}); err != nil {
	}

	jm, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	if _, err := n.js.Publish("wallet.decrease", jm); err != nil {
		return err
	}

	return nil
}
