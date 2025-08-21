package services

import (
	"api-gateway/internal/models"
	"api-gateway/internal/repositories"
	"database/sql"
	"errors"
	"time"
)

type WalletService struct {
	wallets      repositories.WalletRepo
	transactions repositories.TransactionRepo
}

func NewWalletService(c repositories.WalletRepo, t repositories.TransactionRepo) *WalletService {
	return &WalletService{
		wallets:      c,
		transactions: t,
	}
}

func (s *WalletService) GetBalanceByCompanyID(companyId int64) (int64, error) {
	wallet, err := s.wallets.GetWalletByCompanyID(companyId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, repositories.ErrWalletNotFound
		}
		return 0, err
	}

	if wallet == nil {
		return 0, repositories.ErrWalletNotFound
	}

	return wallet.Balance, nil
}

func (s *WalletService) IncreaseWalletBalance(companyId int64, amount int64, refType string, refID int64, idempotencyKey string) error {
	wallet, err := s.wallets.GetWalletByCompanyID(companyId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if wallet == nil {
		wallet = &models.Wallet{
			CompanyId: companyId,
			Balance:   0,
			Version:   0,
			UpdatedAt: time.Now(),
		}
		if err := s.wallets.CreateWallet(wallet); err != nil {
			return err
		}
	}

	wallet.Balance += amount
	wallet.Version++
	wallet.UpdatedAt = time.Now()
	if err = s.wallets.UpdateWallet(wallet); err != nil {
		return err
	}

	trx := &models.Transaction{
		CompanyId:      companyId,
		Amount:         amount,
		Action:         "CREDIT",
		RefType:        refType,
		RefId:          refID,
		IdempotencyKey: idempotencyKey,
		CreatedAt:      time.Now(),
	}
	if err = s.transactions.CreateTransaction(trx); err != nil {
		return err
	}

	return nil
}

func (s *WalletService) DecreaseWalletBalance(companyId int64, amount int64, refType string, refID int64, idempotencyKey string) error {
	wallet, err := s.wallets.GetWalletByCompanyID(companyId)
	if err != nil {
		return err
	}

	if wallet == nil || wallet.Balance < amount {
		return repositories.ErrNoEnoughBalance
	}

	wallet.Balance -= amount
	wallet.Version++
	wallet.UpdatedAt = time.Now()
	if err = s.wallets.UpdateWallet(wallet); err != nil {
		return err
	}

	trx := &models.Transaction{
		CompanyId:      companyId,
		Amount:         amount,
		Action:         "DEBIT",
		RefType:        refType,
		RefId:          refID,
		IdempotencyKey: idempotencyKey,
		CreatedAt:      time.Now(),
	}
	if err := s.transactions.CreateTransaction(trx); err != nil {
		return err
	}

	return nil
}
