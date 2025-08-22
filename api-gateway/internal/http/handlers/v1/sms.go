package v1

import (
	"api-gateway/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

var smsSvc *services.SmsService

type SendSMSRequest struct {
	Message   string `json:"message"`
	Receiver  string `json:"receiver"`
	CompanyID int64  `json:"company_id"`
}

func SendSMS(c echo.Context) (err error) {
	req := new(SendSMSRequest)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	err = services.SmsSrv.SendSMS(c.Request().Context(), req.CompanyID, req.Message, req.Receiver)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
