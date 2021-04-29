package service

import (
	"github.com/niluwats/covid_log/domain"
	"github.com/niluwats/covid_log/dto"
	"github.com/niluwats/covid_log/errs"
)

type VisitorService interface {
	GetAllVisitors() ([]domain.Visitor, *errs.AppError)
	GetVisitor(id string) (*dto.NewVisitorResponse, *errs.AppError)
	NewVisitor(dto.NewVisitorRequest) (*dto.NewVisitorResponse, *errs.AppError)
	EditVisitor(string, dto.NewVisitorRequest) (*dto.NewVisitorResponse, *errs.AppError)
	RemoveVisitor(id string) *errs.AppError
}
type DefaultVisitorService struct {
	repo domain.VisitorRepository
}

func (s DefaultVisitorService) GetAllVisitors() ([]domain.Visitor, *errs.AppError) {
	return s.repo.FindAllVisitors()
}

func (s DefaultVisitorService) GetVisitor(id string) (*dto.NewVisitorResponse, *errs.AppError) {
	v, err := s.repo.FindByNIC(id)
	if err != nil {
		return nil, err
	}
	res := v.ToDto()
	return &res, nil
}

func (s DefaultVisitorService) NewVisitor(req dto.NewVisitorRequest) (*dto.NewVisitorResponse, *errs.AppError) {
	visitor := domain.Visitor{
		NIC:         req.NIC,
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		AddressCity: req.AddressCity,
		ContactNo:   req.ContactNo,
	}
	nic := visitor.NIC
	newVisitor, err := s.repo.SaveVisitor(nic, visitor)
	if err != nil {
		return nil, err
	}
	resp := newVisitor.ToDto()
	return &resp, nil
}
func (s DefaultVisitorService) EditVisitor(id string, req dto.NewVisitorRequest) (*dto.NewVisitorResponse, *errs.AppError) {
	newVisitor := domain.Visitor{
		FirstName:   req.FirstName,
		LastName:    req.LastName,
		AddressCity: req.AddressCity,
		ContactNo:   req.ContactNo,
	}
	updatedNewVisitor, err := s.repo.UpdateVisitor(id, newVisitor)
	if err != nil {
		return nil, err
	}
	resp := updatedNewVisitor.ToDto()
	return &resp, nil
}
func (s DefaultVisitorService) RemoveVisitor(id string) *errs.AppError {
	return s.repo.DeleteVisitor(id)
}
func NewVisitorService(repository domain.VisitorRepository) DefaultVisitorService {
	return DefaultVisitorService{repository}
}
 