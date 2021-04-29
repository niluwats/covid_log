package domain

import (
	"github.com/niluwats/covid_log/dto"
	"github.com/niluwats/covid_log/errs"
)

type Company struct {
	CompanyId   string `db:"company_id" json:"company_id"`
	CompanyName string `db:"name" json:"company_name"`
}

func (c Company) ToDto() dto.NewCompanyResponse {
	return dto.NewCompanyResponse{
		CompanyId:   c.CompanyId,
		CompanyName: c.CompanyName,
	}
}

type CompanyRepository interface {
	SaveCompany(string, Company) (*Company, *errs.AppError)
	UpdateCompany(string, Company) (*Company, *errs.AppError)
	DeleteCompany(string) *errs.AppError
	FindById(id string) (*Company, *errs.AppError)
	FindAll() ([]Company, *errs.AppError)
}
