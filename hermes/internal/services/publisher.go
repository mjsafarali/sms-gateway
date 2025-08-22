package services

import (
	"context"
	"time"
)

type Msg struct {
	CompanyID int64
	Content   string
	Receiver  string
	Timestamp time.Time
	Price     int64
}

type Publisher interface {
	Publish(ctx context.Context, msg Msg) error
}
