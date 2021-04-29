package domain

import (
	"github.com/niluwats/covid_log/dto"
	"github.com/niluwats/covid_log/errs"
)

type Log struct {
	LogId     string `db:"log_id" json:"log_id"`
	NIC       string `db:"nic" json:"nic"`
	CompanyId string `db:"company_id" json:"company_id"`
	Date      string `db:"date" json:"date"`
	LogTime   string `db:"log_time" json:"log_time"`
}

func (l Log) ToDto() dto.NewLogResponse {
	return dto.NewLogResponse{
		NIC:     l.NIC,
		LogTime: l.LogTime,
		Date:    l.Date,
	}
}

type LogRepository interface {
	SaveLog(id string, l Log) (*Log, *errs.AppError)
	FindByDate(id string, date string) ([]dto.NewLogResponseByDate, *errs.AppError)
	FindByVisitorId(nic string) ([]dto.NewLogResponseByNic, *errs.AppError)
	FindAll(id string) ([]dto.NewLogResponseAll, *errs.AppError)
}
