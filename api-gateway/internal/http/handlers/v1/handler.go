package v1

import (
	"api-gateway/internal/repositories"
	"api-gateway/internal/services"
	"github.com/labstack/echo/v4"
)

func Index(c echo.Context) (err error) {
	return c.JSON(200, map[string]string{
		"message": "ok",
	})
}

var svc *services.CompanyService

func CompaniesIndex(c echo.Context) (err error) {
	svc = services.NewCompanyService(repositories.Companies)
	companies, err := svc.GetAllCompanies(c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, companies)
}
