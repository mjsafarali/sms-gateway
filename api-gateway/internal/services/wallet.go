package services

import (
	"api-gateway/internal/config"
	"encoding/json"
	"fmt"
	"net/http"
)

type Wallet struct {
	Addr   string
	Client *http.Client
}

func NewWalletService() *Wallet {
	cfg := config.Cfg.Wallet

	return &Wallet{
		Addr: fmt.Sprintf("%s:%s", cfg.Host, cfg.Port),
		Client: &http.Client{
			Timeout: cfg.Timeout,
		},
	}
}

type BalanceResponse struct {
	Balance int64 `json:"balance"`
}

func (w *Wallet) GetBalance(companyId int64) (int64, error) {
	//TODO: circuit breaker

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
}
