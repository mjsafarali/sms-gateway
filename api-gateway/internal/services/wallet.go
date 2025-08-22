package services

import (
	"api-gateway/internal/config"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/sony/gobreaker"
)

type Wallet struct {
	Addr    string
	Client  *http.Client
	breaker *gobreaker.CircuitBreaker
}

func NewWalletService() *Wallet {
	cfg := config.Cfg.Wallet

	st := gobreaker.Settings{
		Name:        "WalletServiceBreaker",
		MaxRequests: 3,
		Timeout:     5 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			failRatio := float64(counts.TotalFailures) / float64(counts.Requests)
			return counts.Requests >= 5 && failRatio >= 0.5
		},
	}

	return &Wallet{
		Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Client: &http.Client{
			Timeout: cfg.Timeout,
		},
		breaker: gobreaker.NewCircuitBreaker(st),
	}
}

type BalanceResponse struct {
	Balance int64 `json:"balance"`
}

func (w *Wallet) GetBalance(companyId int64) (int64, error) {
	result, err := w.breaker.Execute(func() (interface{}, error) {
		res, err := w.Client.Get(w.Addr + "/api/v1/balance/" + fmt.Sprint(companyId))
		if err != nil {
			return 0, err
		}

		if res.StatusCode != http.StatusOK {
			return 0, fmt.Errorf("failed to get balance: %s", res.Status)
		}

		var balance BalanceResponse
		if err = json.NewDecoder(res.Body).Decode(&balance); err != nil {
			return 0, fmt.Errorf("failed to decode balance response: %w", err)
		}
		defer res.Body.Close()

		return balance.Balance, nil
	})

	if err != nil {
		return 0, err
	}

	return result.(int64), nil
}
