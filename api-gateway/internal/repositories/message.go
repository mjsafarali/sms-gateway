package repositories

import "github.com/jmoiron/sqlx"

type MessagesRepo struct {
	db *sqlx.DB
}

var (
	insertMessageQuery = `INSERT INTO messages (company_id, receiver, content, created_at) VALUES (?, ?, ?, NOW())`
)

func NewMessagesRepository(db *sqlx.DB) *MessagesRepo {
	return &MessagesRepo{
		db: db,
	}
}

func (r *MessagesRepo) CreateMessage(companyID int64, receiver string, content string) error {
	if _, err := r.db.Exec(insertMessageQuery, companyID, receiver, content); err != nil {
		return err
	}

	return nil
}
