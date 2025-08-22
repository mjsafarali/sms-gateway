package services

import (
	"api-gateway/internal/config"
	"api-gateway/internal/repositories"
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"time"
)

type SmsService struct {
	pub Publisher
}

func NewSmsService(pub Publisher) *SmsService {
	return &SmsService{
		pub: pub,
	}
}

func (s *SmsService) SendSMS(ctx context.Context, companyID int64, message, receiver string) error {
	var balance int64
	price := config.Cfg.SMS.EachMessagePrice

	balance, err := s.getBalance(ctx, companyID)
	if err != nil {
		return err
	}

	if balance < price {
		return repositories.ErrInsufficient
	}

	if _, err = repositories.RedisRepository.DecrByNoNegative(ctx, companyID, price); err != nil {
		return err
	}

	msg := Msg{
		CompanyID: companyID,
		Content:   message,
		Receiver:  receiver,
		timestamp: time.Now(),
		Price:     price,
	}
	if err = s.pub.Publish(ctx, msg); err != nil {
		if _, err = repositories.RedisRepository.IncrBy(ctx, companyID, price); err != nil {
			return err
		}
		return err
	}

	return nil
}

func (s *SmsService) getBalance(ctx context.Context, companyID int64) (int64, error) {

	balance, err := repositories.RedisRepository.GetInt64(ctx, companyID)
	if err != nil && !errors.Is(err, redis.Nil) {
		return 0, err
	}

	if errors.Is(err, redis.Nil) {
		dbBalance, err := Wallets.GetBalance(companyID)
		if err != nil {
			return 0, err
		}

		ok, _ := repositories.RedisRepository.SetNX(ctx, companyID, dbBalance)
		if !ok {
			balance, _ = repositories.RedisRepository.GetInt64(ctx, companyID)
		} else {
			balance = dbBalance
		}
	}

	return balance, nil
}
