package services

import "context"

var (
	SmsSrv  SmsService
	Wallets WalletService
	NatsSrv NatsService
)

type SmsService interface {
	SendSMS(ctx context.Context, companyID int64, message string, receiver string) error
}

type WalletService interface {
	GetBalance(companyId int64) (int64, error)
}

type NatsService interface {
	AddStream(name string, subjects []string) error
	Publish(subject string, data []byte) error
}
