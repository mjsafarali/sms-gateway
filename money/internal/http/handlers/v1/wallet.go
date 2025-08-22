package v1

import (
	"errors"
	"github.com/labstack/echo/v4"
	"money/internal/repositories"
	"money/internal/services"
	"strconv"
)

// WalletBalance handles the request to get a company's wallet balance.
func WalletBalance(c echo.Context) (err error) {
	companyID := c.Param("company_id")
	if companyID == "" {
		return c.JSON(400, map[string]string{"error": "company_id is required"})
	}

	cid, err := strconv.ParseInt(companyID, 10, 64)
	if err != nil {
		return c.JSON(400, map[string]string{"error": "invalid company_id"})
	}

	balance, err := services.WalletServiceInstance.GetBalanceByCompanyID(cid)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "internal server error"})
	}

	return c.JSON(200, map[string]interface{}{
		"balance": balance,
	})
}

type IncreaseWalletRequest struct {
	CompanyID int64  `json:"company_id" validate:"required"`
	Amount    int64  `json:"amount" validate:"required"`
	Action    string `json:"action" validate:"required,oneof=CREDIT DEBIT"`
}

// WalletApply Apply handles the request to increase or decrease a company's wallet balance.
func WalletApply(c echo.Context) (err error) {
	req := new(IncreaseWalletRequest)
	if err = c.Bind(req); err != nil {
		return c.JSON(400, map[string]string{"error": "invalid request"})
	}

	if err = c.Validate(req); err != nil {
		return c.JSON(400, map[string]string{"error": "validation failed"})
	}

	if req.Action == "CREDIT" {
		if err = services.WalletServiceInstance.IncreaseWalletBalance(req.CompanyID, req.Amount); err != nil {
			return c.JSON(500, map[string]string{"error": "internal server error"})
		}
	} else {
		if err = services.WalletServiceInstance.DecreaseWalletBalance(req.CompanyID, req.Amount); err != nil {
			if errors.Is(err, repositories.ErrNoEnoughBalance) {
				return c.JSON(422, map[string]string{"error": "not enough balance"})
			}
			return c.JSON(500, map[string]string{"error": "internal server error"})
		}
	}

	return c.JSON(200, map[string]string{
		"message": "success",
	})
}
