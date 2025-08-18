package repositories

import (
	"api-gateway/internal/models"
	"context"
	"errors"
)

var (
	ErrNoQueryResult = errors.New("no query result")
	Companies        CompanyRepo
)

type CompanyRepo interface {
	GetAllCompanies(ctx context.Context) ([]models.Company, error)
}
