package repositories

import (
	"api-gateway/internal/models"
	"github.com/jmoiron/sqlx"
)

type MysqlWallet struct {
	db *sqlx.DB
}

var (
	insertWalletQuery         = `INSERT INTO wallets (company_id, balance, version, updated_at) VALUES (?, ?, ?, ?)`
	getWalletByCompanyIDQuery = `SELECT company_id, balance, version, updated_at FROM wallets WHERE company_id = ?`
	updateWalletQuery         = `UPDATE wallets SET balance = ?, version = ?, updated_at = ? WHERE company_id = ?`
)

func NewMysqlWallet(db *sqlx.DB) *MysqlWallet {
	return &MysqlWallet{
		db: db,
	}
}

func (w *MysqlWallet) CreateWallet(wallet *models.Wallet) error {
	if _, err := w.db.Exec(insertWalletQuery, wallet.CompanyId, wallet.Balance, wallet.Version, wallet.UpdatedAt); err != nil {
		return err
	}

	return nil
}

func (w *MysqlWallet) GetWalletByCompanyID(id int64) (*models.Wallet, error) {
	wallet := &models.Wallet{}
	if err := w.db.QueryRow(getWalletByCompanyIDQuery, id).Scan(&wallet.CompanyId, &wallet.Balance, &wallet.Version, &wallet.UpdatedAt); err != nil {
		return nil, err
	}

	return wallet, nil
}

func (w *MysqlWallet) UpdateWallet(wallet *models.Wallet) error {
	if _, err := w.db.Exec(updateWalletQuery, wallet.Balance, wallet.Version, wallet.UpdatedAt, wallet.CompanyId); err != nil {
		return err
	}

	return nil
}
