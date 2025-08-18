package models

import "time"

type Wallet struct {
	CompanyId int64     `json:"company_id"`
	Balance   int64     `json:"balance"`
	Version   int64     `json:"version"`
	UpdatedAt time.Time `json:"updated_at"`
}
