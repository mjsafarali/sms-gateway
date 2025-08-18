package models

import "time"

type Company struct {
	Id          int64     `json:"id"`
	Name        string    `json:"name"`
	PricePerSms int64     `json:"price_per_sms"`
	DailyQuota  int64     `json:"daily_quota"`
	RpsLimit    int       `json:"rps_limit"`
	IsActive    bool      `json:"is_active"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
