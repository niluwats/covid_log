package domain

import (
	"github.com/niluwats/covid_log/dto"
	"github.com/niluwats/covid_log/errs"
)

type Visitor struct {
	NIC         string `db:"nic" json:"nic"`
	FirstName   string `db:"fname" json:"first_name"`
	LastName    string `db:"lname" json:"last_name"`
	AddressCity string `db:"address_city" json:"address_city"`
	ContactNo   string `db:"contact_no" json:"contact_no"`
}

func (v Visitor) ToDto() dto.NewVisitorResponse {
	return dto.NewVisitorResponse{
		NIC:         v.NIC,
		FirstName:   v.FirstName,
		LastName:    v.LastName,
		AddressCity: v.AddressCity,
		ContactNo:   v.ContactNo,
	}
}

type VisitorRepository interface {
	SaveVisitor(string,Visitor) (*Visitor, *errs.AppError)
	UpdateVisitor(string, Visitor) (*Visitor, *errs.AppError)
	DeleteVisitor(string) *errs.AppError
	FindByNIC(id string) (*Visitor, *errs.AppError)
	FindAllVisitors() ([]Visitor, *errs.AppError)
}
