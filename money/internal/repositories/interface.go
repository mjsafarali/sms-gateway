package repositories

import (
	"api-gateway/internal/models"
	"errors"
)

var (
	ErrNoQueryResult   = errors.New("no query result")
	ErrNoEnoughBalance = errors.New("not enough balance")
	ErrWalletNotFound  = errors.New("wallet not found")
	ErrTrxExists       = errors.New("transaction already exists")

	Wallets      WalletRepo
	Transactions TransactionRepo
)

type WalletRepo interface {
	CreateWallet(wallet *models.Wallet) error
	GetWalletByCompanyID(id int64) (*models.Wallet, error)
	UpdateWallet(wallet *models.Wallet) error
}

type TransactionRepo interface {
	CreateTransaction(trx *models.Transaction) error
}
