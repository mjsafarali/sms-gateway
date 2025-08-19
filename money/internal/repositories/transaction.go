package repositories

import (
	"github.com/jmoiron/sqlx"
	"money/internal/models"
)

type MysqlTransaction struct {
	db *sqlx.DB
}

var (
	insertTransactionQuery = `INSERT INTO transactions (company_id, amount, action, ref_type, ref_id, idempotency_key, created_at) VALUES (?, ?, ?, ?, ?, ?, NOW())`
)

func NewMysqlTransaction(db *sqlx.DB) *MysqlTransaction {
	return &MysqlTransaction{
		db: db,
	}
}

func (t *MysqlTransaction) CreateTransaction(trx *models.Transaction) error {
	if _, err := t.db.Exec(insertTransactionQuery, trx.CompanyId, trx.Amount, trx.Action, trx.RefType, trx.RefId, trx.IdempotencyKey); err != nil {
		return err
	}

	return nil
}
