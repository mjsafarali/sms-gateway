package v1

import (
	"api-gateway/internal/services"
	"net/http"

	"github.com/labstack/echo/v4"
)

var smsSvc *services.SmsService

type SendSMSRequest struct {
	Message  string `form:"message"`
	Receiver string `form:"receiver"`
}

func SendSMS(c echo.Context) (err error) {
	form := new(SendSMSRequest)
	if err = c.Bind(form); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
	}

	if smsSvc == nil {
		smsSvc = services.NewSmsService()
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
	})
}
