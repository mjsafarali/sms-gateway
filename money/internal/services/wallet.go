package services

import (
	"money/internal/repositories"
)

type WalletService struct {
	wallets repositories.WalletRepo
}

func NewWalletService(c repositories.WalletRepo) *WalletService {
	return &WalletService{
		wallets: c,
	}
}
