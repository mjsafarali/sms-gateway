package services

import (
	"api-gateway/internal/models"
	"api-gateway/internal/repositories"
	"context"
)

type CompanyService struct {
	companies repositories.CompanyRepo
}

func NewCompanyService(c repositories.CompanyRepo) *CompanyService {
	return &CompanyService{
		companies: c,
	}
}

func (s *CompanyService) GetAllCompanies(ctx context.Context) ([]models.Company, error) {
	companies, err := s.companies.GetAllCompanies(ctx)
	if err != nil {
		return nil, err
	}

	return companies, nil
}
