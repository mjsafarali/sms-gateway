package handlers

import (
	"context"
	"encoding/json"
	"hermes/internal/services"
	"hermes/log"
	"time"

	"github.com/nats-io/nats.go"
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

	walletData := services.Msg{
		CompanyID: data.CompanyID,
		Content:   data.Content,
		Receiver:  data.Receiver,
		Timestamp: time.Now(),
		Price:     data.Price,
	}
	if err := services.PublisherSrv.Publish(context.Background(), walletData); err != nil {
		return
	}
}
