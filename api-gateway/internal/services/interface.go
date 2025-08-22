package services

var (
	SmsSrv  *SmsService
	Wallets WalletService
)

type WalletService interface {
	GetBalance(companyId int64) (int64, error)
}
