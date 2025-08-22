package services

import (
	"context"
	"hermes/log"
)

var SenderSrv Sender

type Sender interface {
	Send(ctx context.Context, to string, message string) error
}

type LogSender struct{}

func NewLogSender() *LogSender {
	return &LogSender{}
}

func (l *LogSender) Send(ctx context.Context, to string, message string) error {
	log.Printf("[LOG SMS] To: %s | Message: %s", to, message)
	return nil
}
