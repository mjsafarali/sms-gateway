package repositories

import (
	"context"
	"github.com/go-redis/redis/v8"
)

type Redis struct {
	rdb *redis.Client
}

func NewRedis(rdb *redis.Client) *Redis {
	return &Redis{
		rdb: rdb,
	}
}

func (r *Redis) GetInt64(ctx context.Context, key string) (int64, error) {
	k := RedisPrefix + key
	val, err := r.rdb.Get(ctx, k).Int64()

	return val, err
}

func (r *Redis) SetNX(ctx context.Context, key string, value interface{}) (bool, error) {
	k := RedisPrefix + key
	ok, err := r.rdb.SetNX(ctx, k, value, 0).Result()

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

func (r *Redis) DecrByNoNegative(ctx context.Context, key string, n int64) (int64, error) {
	k := RedisPrefix + key
	result, err := decrNoNeg.Run(ctx, r.rdb, []string{k}, n).Result()
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

func (r *Redis) IncrBy(ctx context.Context, key string, n int64) (int64, error) {
	k := RedisPrefix + key
	result, err := r.rdb.IncrBy(ctx, k, n).Result()
	if err != nil {
		return 0, err
	}
	return result, nil
}
