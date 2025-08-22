package handlers

import (
	"context"
	"encoding/json"
	"github.com/nats-io/nats.go"
	"hermes/internal/services"
	"hermes/log"
	"time"
)

type Msg struct {
	CompanyID int64
	Content   string
	Receiver  string
	timestamp time.Time
	Price     int64
}

func SendSMSHandler(msg *nats.Msg) {
	var data Msg
	if err := json.Unmarshal(msg.Data, &data); err != nil {
		return
	}

	err := services.SenderSrv.Send(context.Background(), data.Receiver, data.Content)
	if err != nil {
		return
	}

	if err := msg.Ack(); err != nil {
		log.Printf("failed to ack message: %v", err)
		return
	}

}
