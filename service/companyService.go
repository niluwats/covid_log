package service

import (
	"github.com/niluwats/covid_log/domain"
	"github.com/niluwats/covid_log/dto"
	"github.com/niluwats/covid_log/errs"
)

type CompanyService interface {
	GetAllComanpies() ([]domain.Company, *errs.AppError)
	GetCompany(id string) (*dto.NewCompanyResponse, *errs.AppError)
	NewCompany(dto.NewCompanyRequest) (*dto.NewCompanyResponse, *errs.AppError)
	EditCompany(string, dto.NewCompanyRequest) (*dto.NewCompanyResponse, *errs.AppError)
	RemoveCompany(id string) *errs.AppError
}
type DefaultCompanyService struct {
	repo domain.CompanyRepository
}

func (s DefaultCompanyService) GetAllComanpies() ([]domain.Company, *errs.AppError) {
	return s.repo.FindAll()
}
func (s DefaultCompanyService) GetCompany(id string) (*dto.NewCompanyResponse, *errs.AppError) {
	c, err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}
	res := c.ToDto()
	return &res, nil
}
func (s DefaultCompanyService) NewCompany(req dto.NewCompanyRequest) (*dto.NewCompanyResponse, *errs.AppError) {
	comp := domain.Company{
		CompanyId:   req.CompanyId,
		CompanyName: req.CompanyName,
	}
	id := comp.CompanyId
	newComp, err := s.repo.SaveCompany(id, comp)
	if err != nil {
		return nil, err
	}
	resp := newComp.ToDto()
	return &resp, nil
}
func (s DefaultCompanyService) EditCompany(id string, req dto.NewCompanyRequest) (*dto.NewCompanyResponse, *errs.AppError) {
	newComp := domain.Company{
		// CompanyId:   req.CompanyId,
		CompanyName: req.CompanyName,
	}
	updatedNewComp, err := s.repo.UpdateCompany(id, newComp)
	if err != nil {
		return nil, err
	}
	resp := updatedNewComp.ToDto()
	return &resp, nil
}
func (s DefaultCompanyService) RemoveCompany(id string) *errs.AppError {
	return s.repo.DeleteCompany(id)
}
func NewCompanyService(repo domain.CompanyRepository) DefaultCompanyService {
	return DefaultCompanyService{repo}
}
