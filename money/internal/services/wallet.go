package services

import (
	"api-gateway/internal/models"
	"api-gateway/internal/repositories"
	"time"
)

var WalletServiceInstance *WalletService

type WalletService struct {
	transactions repositories.TransactionRepo
}

func NewWalletService(t repositories.TransactionRepo) *WalletService {
	return &WalletService{
		transactions: t,
	}
}

func (s *WalletService) GetBalanceByCompanyID(companyId int64) (int64, error) {
	balance, err := s.transactions.GetLatestBalance(companyId)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (s *WalletService) IncreaseWalletBalance(companyId int64, amount int64) error {
	trx := &models.Transaction{
		CompanyId: companyId,
		Amount:    amount,
		Action:    "CREDIT",
		CreatedAt: time.Now(),
	}
	if err := s.transactions.CreateTransaction(trx); err != nil {
		return err
	}

	return nil
}

func (s *WalletService) DecreaseWalletBalance(companyId int64, amount int64) error {
	trx := &models.Transaction{
		CompanyId: companyId,
		Amount:    amount,
		Action:    "DEBIT",
		CreatedAt: time.Now(),
	}
	if err := s.transactions.CreateTransaction(trx); err != nil {
		return err
	}

	return nil
}
