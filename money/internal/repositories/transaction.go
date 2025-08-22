package repositories

import (
	"api-gateway/internal/models"
	"database/sql"
	"errors"
	"github.com/jmoiron/sqlx"
)

type MysqlTransaction struct {
	db *sqlx.DB
}

var (
	insertTransactionQuery = `INSERT INTO transactions (company_id, amount, action, balance, created_at) VALUES (?, ?, ?, ?, NOW())`
	getBalanceQuery        = `SELECT balance FROM transactions WHERE company_id = ? ORDER BY created_at DESC LIMIT 1`
)

func NewMysqlTransaction(db *sqlx.DB) *MysqlTransaction {
	return &MysqlTransaction{
		db: db,
	}
}

func (t *MysqlTransaction) CreateTransaction(trx *models.Transaction) error {
	balance, err := t.GetLatestBalance(trx.CompanyId)
	if err != nil {
		return err
	}

	if trx.Action == "CREDIT" {
		balance += trx.Amount
	} else {
		if balance < trx.Amount {
			return ErrNoEnoughBalance
		}
		balance -= trx.Amount
	}

	if _, err := t.db.Exec(insertTransactionQuery, trx.CompanyId, trx.Amount, trx.Action, balance); err != nil {
		return err
	}

	return nil
}

func (t *MysqlTransaction) GetLatestBalance(companyID int64) (int64, error) {
	var balance int64
	if err := t.db.Get(&balance, getBalanceQuery, companyID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, nil
		}
		return 0, err
	}
	return balance, nil
}
