package repositories

import (
	"api-gateway/internal/models"
	"context"
	"github.com/jmoiron/sqlx"
)

type MysqlCompanyRepo struct {
	db *sqlx.DB
}

func NewMysqlCompanyRepo(db *sqlx.DB) *MysqlCompanyRepo {
	return &MysqlCompanyRepo{
		db: db,
	}
}

func (r *MysqlCompanyRepo) GetAllCompanies(ctx context.Context) ([]models.Company, error) {
	rows, err := r.db.QueryContext(ctx, `SELECT * FROM companies ORDER BY id DESC`)
	if err != nil {
		return nil, err
	}

	var companies []models.Company

	for rows.Next() {
		var company models.Company
		if err = rows.Scan(
			&company.Id,
			&company.Name,
			&company.PricePerSms,
			&company.DailyQuota,
			&company.RpsLimit,
			&company.IsActive,
			&company.CreatedAt,
			&company.UpdatedAt,
		); err != nil {
			return nil, err
		}
		companies = append(companies, company)
	}
	return companies, nil
}
