package repositories

import (
	"github.com/jmoiron/sqlx"
)

type MysqlWallet struct {
	db *sqlx.DB
}

func NewMysqlWallet(db *sqlx.DB) *MysqlWallet {
	return &MysqlWallet{
		db: db,
	}
}
