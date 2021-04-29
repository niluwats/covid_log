package domain

import (
	"github.com/jmoiron/sqlx"
	"github.com/niluwats/covid_log/dto"
	"github.com/niluwats/covid_log/errs"
	"github.com/niluwats/covid_log/logger"
)

type LogRepositoryDb struct {
	client *sqlx.DB
}

func (d LogRepositoryDb) SaveLog(id string, l Log) (*Log, *errs.AppError) {
	_, err := d.client.Exec("INSERT INTO log(log_time,date,nic,company_id) VALUES(?,?,?,?)", l.LogTime, l.Date, l.NIC, id)
	if err != nil {
		logger.Error("error while inserting a new log" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected DB error")
	}
	return &l, nil
}
func (d LogRepositoryDb) FindAll(id string) ([]dto.NewLogResponseAll, *errs.AppError) {
	findAllSql := "SELECT date,log_time,v.nic,fname,lname,address_city,contact_no FROM log l INNER JOIN visitors v ON v.nic=l.nic WHERE company_id=?"
	logs := make([]dto.NewLogResponseAll, 0)
	err := d.client.Select(&logs, findAllSql, id)
	if err != nil {
		logger.Error("error while querying all the logs " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected DB error")
	}
	return logs, nil
}
func (d LogRepositoryDb) FindByDate(id string, date string) ([]dto.NewLogResponseByDate, *errs.AppError) {
	findSql := "SELECT log_time,date,v.nic,fname,lname,address_city,contact_no FROM log l inner JOIN visitors v ON l.nic=v.nic WHERE company_id=? and date=?"
	logs := make([]dto.NewLogResponseByDate, 0)
	err := d.client.Select(&logs, findSql, id, date)
	if err != nil {
		logger.Error("error while querying the logs by date " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected DB error")
	}
	return logs, nil
}
func (d LogRepositoryDb) FindByVisitorId(nic string) ([]dto.NewLogResponseByNic, *errs.AppError) {
	findSql := "SELECT date,log_time,c.company_id,name FROM log l INNER JOIN company c on l.company_id=c.company_id WHERE nic=?"
	logs := make([]dto.NewLogResponseByNic, 0)
	err := d.client.Select(&logs, findSql, nic)
	if err != nil {
		logger.Error("error while querying logs by NIC " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected DB error")
	}
	return logs, nil
}
func NewLogRepositoryDb(dbClient *sqlx.DB) LogRepositoryDb {
	return LogRepositoryDb{dbClient}
}
