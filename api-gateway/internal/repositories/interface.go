package repositories

import (
	"context"
	"errors"
)

var (
	RedisPrefix = "api-gateway:"

	ErrNoQueryResult = errors.New("no query result")
	ErrInsufficient  = errors.New("decrement would go below zero")

	RedisRepository RedisRepo
)

type RedisRepo interface {
	GetInt64(ctx context.Context, key string) (int64, error)
	SetNX(ctx context.Context, key string, value interface{}) (bool, error)
	DecrByNoNegative(ctx context.Context, key string, n int64) (int64, error)
	IncrBy(ctx context.Context, key string, n int64) (int64, error)
}
