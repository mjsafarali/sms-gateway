package models

import "time"

type Transaction struct {
	Id             int64     `json:"id"`
	CompanyId      int64     `json:"company_id"`
	Amount         int64     `json:"balance"`
	Action         string    `json:"action"`
	RefType        string    `json:"ref_type"`
	RefId          int64     `json:"ref_id"`
	IdempotencyKey string    `json:"idempotency_key"`
	CreatedAt      time.Time `json:"created_at"`
}
