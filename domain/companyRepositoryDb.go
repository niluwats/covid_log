package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/niluwats/covid_log/errs"
	"github.com/niluwats/covid_log/logger"
)

type CompanyRepositoryDb struct {
	client *sqlx.DB
}

func (d CompanyRepositoryDb) SaveCompany(id string, c Company) (*Company, *errs.AppError) {
	findIfSql := "SELECT * FROM company WHERE company_id=?"
	err0 := d.client.Get(&c, findIfSql, id)
	if err0 != nil {
		_, err := d.client.Exec("INSERT INTO company values(?,?)", c.CompanyId, c.CompanyName)
		if err != nil {
			logger.Error("error while saving a new company " + err.Error())
			return nil, errs.NewUnexpectedError("unexpexted DB error")
		}
		return &c, nil
	} else {
		logger.Error("Duplicate entry")
		return nil, errs.NewUnexpectedError("this companyID already exists under this ID ")
	}
}

func (d CompanyRepositoryDb) FindAll() ([]Company, *errs.AppError) {
	var err error
	comps := make([]Company, 0)
	findAllSql := "SELECT * FROM company"
	err = d.client.Select(&comps, findAllSql)
	if err != nil {
		logger.Error("error while querying companies" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected DB error")
	}
	return comps, nil
}
func (d CompanyRepositoryDb) FindById(id string) (*Company, *errs.AppError) {
	var c Company
	var err error
	findByIdSql := "SELECT * FROM company WHERE company_id=?"
	err = d.client.Get(&c, findByIdSql, id)
	if err != nil {
		logger.Error("error while querying the company by id " + err.Error())
		return nil, errs.NewUnexpectedError("No such company under this ID")
	}
	return &c, nil
}
func (d CompanyRepositoryDb) UpdateCompany(id string, c Company) (*Company, *errs.AppError) {
	_, err := d.client.Exec("UPDATE company SET name=? WHERE company_id=?", c.CompanyName, id)
	if err != nil {
		logger.Error("error while updating company name " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected DB error")
	}
	return &c, nil
}
func (d CompanyRepositoryDb) DeleteCompany(id string) *errs.AppError {
	_, err := d.client.Exec("DELETE FROM company WHERE company_id=?", id)
	if err != nil {
		logger.Error("error while deleting company " + err.Error())
		return errs.NewUnexpectedError("unexpected DB error")
	}
	return nil
}
func NewCompanyRepositoryDb(dbClient *sqlx.DB) CompanyRepositoryDb {
	return CompanyRepositoryDb{dbClient}
}
