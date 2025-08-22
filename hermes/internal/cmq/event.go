package cmq

import (
	"github.com/nats-io/nats.go"
	"hermes/internal/app"
	"hermes/internal/cmq/handlers"
	"hermes/log"
)

const (
	QGroup = "SMS"
	smsSub = "sms.send"
)

// Consumer is the cmq consumer struct.
//
// It holds a sync.WaitGroup to handle goroutines.
type Consumer struct {
	subscription *nats.Subscription
}

// NewConsumer is a function for creating a Consumer object.
func NewConsumer() *Consumer {
	return &Consumer{}
}

func (c *Consumer) Start() *Consumer {
	sub, err := app.A.NatsJS.QueueSubscribe(
		smsSub,
		QGroup,
		handlers.SendSMSHandler,
		nats.DeliverNew(),
	)
	if err != nil {
		log.Fatalf("error while subscribing to nats, error: %+v", err)
	}

	c.subscription = sub
	return c
}

func (c *Consumer) WaitForSignals(shutdownRequest chan struct{}) (shutdownReady chan struct{}) {
	shutdownReady = make(chan struct{})
	go func() {
		<-shutdownRequest

		if c.subscription != nil {
			if err := c.subscription.Unsubscribe(); err != nil {
				log.Debugf("error while unsubscribing from cmq error : %+v", err)
			}
		}

		shutdownReady <- struct{}{}
	}()

	return
}
