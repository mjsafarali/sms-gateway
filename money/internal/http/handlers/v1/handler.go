package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"money/internal/repositories"
	"money/internal/services"
	"money/log"
	"strconv"
)

var svc *services.WalletService

func Balance(c echo.Context) (err error) {
	companyID := c.Param("company_id")
	if companyID == "" {
		return c.JSON(400, map[string]string{"error": "company_id is required"})
	}

	cid, err := strconv.ParseInt(companyID, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "invalid company_id"})
	}

	if svc == nil {
		svc = services.NewWalletService(repositories.Wallets, repositories.Transactions)
	}

	balance, err := svc.GetBalanceByCompanyID(cid)
	if err != nil {
		if errors.Is(err, repositories.ErrWalletNotFound) {
			return c.JSON(404, map[string]string{"error": "wallet not found"})
		}
		return c.JSON(500, map[string]string{"error": "internal server error"})
	}

	return c.JSON(200, map[string]interface{}{
		"balance": balance,
	})
}

type IncreaseWalletRequest struct {
	CompanyID      int64  `json:"company_id" validate:"required"`
	Amount         int64  `json:"amount" validate:"required"`
	Action         string `json:"action" validate:"required,oneof=CREDIT DEBIT"`
	RefType        string `json:"ref_type" validate:"required"`
	RefID          int64  `json:"ref_id" validate:"required"`
	IdempotencyKey string `json:"idempotency_key" validate:"required"`
}

// Apply handles the request to increase or decrease a company's wallet balance.
func Apply(c echo.Context) (err error) {
	req := new(IncreaseWalletRequest)
	if err = c.Bind(req); err != nil {
		return c.JSON(400, map[string]string{"error": "invalid request"})
	}

	if err = c.Validate(req); err != nil {
		log.Info(err.Error())
		return c.JSON(400, map[string]string{"error": "validation failed"})
	}

	if svc == nil {
		svc = services.NewWalletService(repositories.Wallets, repositories.Transactions)
	}

	if req.Action == "CREDIT" {
		if err = svc.IncreaseWalletBalance(req.CompanyID, req.Amount, req.RefType, req.RefID, req.IdempotencyKey); err != nil {
			if errors.Is(err, repositories.ErrTrxExists) {
				return c.JSON(422, map[string]string{"error": "transaction already exists"})
			}
			return c.JSON(500, map[string]string{"error": "internal server error"})
		}
	} else if req.Action == "DEBIT" {
		if err = svc.DecreaseWalletBalance(req.CompanyID, req.Amount, req.RefType, req.RefID, req.IdempotencyKey); err != nil {
			if errors.Is(err, repositories.ErrNoEnoughBalance) {
				return c.JSON(422, map[string]string{"error": "not enough balance"})
			}
			if errors.Is(err, repositories.ErrTrxExists) {
				return c.JSON(422, map[string]string{"error": "transaction already exists"})
			}
			return c.JSON(500, map[string]string{"error": "internal server error"})
		}
	}

	return c.JSON(200, map[string]string{
		"message": "success",
	})
}
