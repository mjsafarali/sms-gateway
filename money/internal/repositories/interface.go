package repositories

import (
	"errors"
	"money/internal/models"
)

var (
	ErrNoEnoughBalance = errors.New("not enough balance")

	Transactions TransactionRepo
)

type TransactionRepo interface {
	CreateTransaction(trx *models.Transaction) error
	GetLatestBalance(companyID int64) (int64, error)
}
