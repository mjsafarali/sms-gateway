package handlers

import (
	"encoding/json"
	"money/internal/models"
	"money/internal/repositories"
	"money/log"
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

func DecreaseWalletHandler(msg *nats.Msg) {
	var data Msg
	if err := json.Unmarshal(msg.Data, &data); err != nil {
		return
	}

	trx := models.Transaction{
		CompanyId: data.CompanyID,
		Amount:    data.Price,
		Action:    "DEBIT",
	}
	if err := repositories.Transactions.CreateTransaction(&trx); err != nil {
		log.Errorf("failed to create transaction: %v", err)
		return
	}

	if err := msg.Ack(); err != nil {
		log.Printf("failed to ack message: %v", err)
		return
	}
}
