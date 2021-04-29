package domain

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/niluwats/covid_log/errs"
	"github.com/niluwats/covid_log/logger"
)

type VisitorRepositoryDb struct {
	client *sqlx.DB
}

func (d VisitorRepositoryDb) FindAllVisitors() ([]Visitor, *errs.AppError) {
	var err error
	visitors := make([]Visitor, 0)
	findAllSql := "SELECT * FROM visitors"
	err = d.client.Select(&visitors, findAllSql)
	if err != nil {
		logger.Error("error while querying visitors" + err.Error())
		return nil, errs.NewUnexpectedError("unexpected DB error")
	}
	return visitors, nil
}
func (d VisitorRepositoryDb) FindByNIC(id string) (*Visitor, *errs.AppError) {
	var v Visitor
	var err error
	findByIdSql := "SELECT * FROM visitors WHERE nic=?"
	err = d.client.Get(&v, findByIdSql, id)
	if err != nil {
		logger.Error("error while querying the visitor by id " + err.Error())
		return nil, errs.NewUnexpectedError("No such visitor under this NIC")
	}
	return &v, nil
}

func (d VisitorRepositoryDb) SaveVisitor(nic string, v Visitor) (*Visitor, *errs.AppError) {
	findIfSql := "SELECT * FROM visitors WHERE nic=?"
	err0 := d.client.Get(&v, findIfSql, nic)
	if err0 != nil {
		_, err := d.client.Exec("INSERT INTO visitors(nic,fname,lname,address_city,contact_no) values(?,?,?,?,?)", v.NIC, v.FirstName, v.LastName, v.AddressCity, v.ContactNo)
		if err != nil {
			logger.Error("error while inserting a visitor " + err.Error())
			return nil, errs.NewUnexpectedError("unexpexted DB error")
		}
		return &v, nil
	} else {
		logger.Error("Duplicate entry")
		return nil, errs.NewUnexpectedError("this visitor NIC already exists")
	}
}
func (d VisitorRepositoryDb) UpdateVisitor(id string, v Visitor) (*Visitor, *errs.AppError) {
	_, err := d.client.Exec("UPDATE visitors SET fname=?,lname=?,address_city=?,contact_no=? WHERE nic=?", v.FirstName, v.LastName, v.AddressCity, v.ContactNo, id)
	if err != nil {
		logger.Error("error while updating visitor " + err.Error())
		return nil, errs.NewUnexpectedError("unexpected DB error")
	}
	return &v, nil
}

func (d VisitorRepositoryDb) DeleteVisitor(id string) *errs.AppError {
	_, err := d.client.Exec("DELETE FROM visitors WHERE nic=?", id)
	if err != nil {
		logger.Error("error while deleting visitor " + err.Error())
		return errs.NewUnexpectedError("unexpected DB error")
	}
	return nil
}

// func (d VisitorRepositoryDb) IsAvailableNic(req dto.NewVisitorRequest)(*errs.AppError){

// }

func NewVisitorRepositoryDb(dbClient *sqlx.DB) VisitorRepositoryDb {
	return VisitorRepositoryDb{dbClient}
}
