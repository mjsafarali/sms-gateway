package v1

import "github.com/labstack/echo/v4"

func Index(c echo.Context) (err error) {
	return c.JSON(200, map[string]string{
		"message": "ok",
	})
}
