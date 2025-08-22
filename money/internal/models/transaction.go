package models

import "time"

type Transaction struct {
	Id        int64     `json:"id"`
	CompanyId int64     `json:"company_id"`
	Amount    int64     `json:"amount"`
	Action    string    `json:"action"`
	Balance   int64     `json:"balance"`
	CreatedAt time.Time `json:"created_at"`
}
