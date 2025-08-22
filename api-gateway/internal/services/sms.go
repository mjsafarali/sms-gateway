package services

import (
	"api-gateway/internal/config"
	"api-gateway/internal/repositories"
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

type Sms struct {
}

func NewSmsService() *Sms {
	return &Sms{}
}

func (s *Sms) SendSMS(ctx context.Context, companyID int64, message, receiver string) error {
	var balance int64
	price := config.Cfg.SMS.EachMessagePrice
	balanceKey := fmt.Sprintf("company:%s:balance", strconv.Itoa(int(companyID)))

	balance, err := repositories.RedisRepository.GetInt64(ctx, balanceKey)
	if err != nil && !errors.Is(err, redis.Nil) {
		return err
	}

	if errors.Is(err, redis.Nil) {
		dbBalance, err := Wallets.GetBalance(companyID)
		if err != nil {
			return err
		}

		ok, _ := repositories.RedisRepository.SetNX(ctx, balanceKey, dbBalance)
		if !ok {
			balance, _ = repositories.RedisRepository.GetInt64(ctx, balanceKey)
		} else {
			balance = dbBalance
		}
	}

	if balance < price {
		return repositories.ErrInsufficient
	}

	if _, err = repositories.RedisRepository.DecrByNoNegative(ctx, balanceKey, price); err != nil {
		return err
	}

	if err = NatsSrv.AddStream("SMS", []string{"sms.*"}); err != nil {
		if _, err = repositories.RedisRepository.IncrBy(ctx, balanceKey, price); err != nil {
			return err
		}
		return err
	}

	msg := fmt.Sprintf(
		`{"company_id":%d,"message":"%s","receiver":"%s","timestamp":%d,"price":%d}`,
		companyID, message, receiver, time.Now().UnixMilli(), price,
	)
	err = NatsSrv.Publish("sms.send", []byte(msg))
	if err != nil {
		if _, err = repositories.RedisRepository.IncrBy(ctx, balanceKey, price); err != nil {
			return err
		}
		return err
	}

	return nil
}
