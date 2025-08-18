package repositories

import (
	"errors"
)

var (
	ErrNoQueryResult = errors.New("no query result")
	Wallets          WalletRepo
)

type WalletRepo interface {
	//
}
