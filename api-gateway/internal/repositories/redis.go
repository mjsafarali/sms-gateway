package repositories

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

type RedisRepo struct {
	Rdb *redis.Client
}

func NewRedis(rdb *redis.Client) *RedisRepo {
	return &RedisRepo{
		Rdb: rdb,
	}
}

func (r *RedisRepo) GetInt64(ctx context.Context, companyID int64) (int64, error) {
	k := RedisPrefix + key(companyID)
	val, err := r.Rdb.Get(ctx, k).Int64()

	return val, err
}

func (r *RedisRepo) SetNX(ctx context.Context, companyID int64, value interface{}) (bool, error) {
	k := RedisPrefix + key(companyID)
	ok, err := r.Rdb.SetNX(ctx, k, value, 0).Result()

	return ok, err
}

var decrNoNeg = redis.NewScript(`
local balance = redis.call("GET", KEYS[1])
if not balance then
  return -1
end
balance = tonumber(balance)
if balance >= tonumber(ARGV[1]) then
  return redis.call("DECRBY", KEYS[1], ARGV[1])
else
  return -2
end
`)

func (r *RedisRepo) DecrByNoNegative(ctx context.Context, companyID int64, n int64) (int64, error) {
	k := RedisPrefix + key(companyID)
	result, err := decrNoNeg.Run(ctx, r.Rdb, []string{k}, n).Result()
	if err != nil {
		return 0, err
	}

	if result == -1 {
		return 0, ErrNoQueryResult
	} else if result == -2 {
		return 0, ErrInsufficient
	}

	return result.(int64), nil
}

func (r *RedisRepo) IncrBy(ctx context.Context, companyID int64, n int64) (int64, error) {
	k := RedisPrefix + key(companyID)
	result, err := r.Rdb.IncrBy(ctx, k, n).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}

func key(companyID int64) string {
	return fmt.Sprintf("company:%d:balance", companyID)
}
