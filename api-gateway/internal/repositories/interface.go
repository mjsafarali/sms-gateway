package repositories

import (
	"errors"
)

var (
	RedisPrefix = "api-gateway:"

	ErrNoQueryResult = errors.New("no query result")
	ErrInsufficient  = errors.New("decrement would go below zero")

	RedisRepository *RedisRepo
)
