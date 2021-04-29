package service

import (
	"time"

	"github.com/niluwats/covid_log/domain"
	"github.com/niluwats/covid_log/dto"
	"github.com/niluwats/covid_log/errs"
)

type LogService interface {
	NewLog(string, dto.NewLogRequest) (*dto.NewLogResponse, *errs.AppError)
	GetAllLogs(id string) ([]dto.NewLogResponseAll, *errs.AppError)
	GetLogsByVisitorId(id string) ([]dto.NewLogResponseByNic, *errs.AppError)
	GetLogsByDate(id string, date string) ([]dto.NewLogResponseByDate, *errs.AppError)
}

type DefaultLogService struct {
	repo domain.LogRepository
}

func (s DefaultLogService) NewLog(id string, req dto.NewLogRequest) (*dto.NewLogResponse, *errs.AppError) {
	curTime := time.Now().Format("15:04:05")
	date := time.Now().Format("2006-01-02")
	log := domain.Log{
		NIC:     req.NIC,
		LogTime: curTime,
		Date:    date,
	}
	newLog, err := s.repo.SaveLog(id, log)
	if err != nil {
		return nil, err
	}
	resp := newLog.ToDto()
	return &resp, nil
}
func (s DefaultLogService) GetAllLogs(id string) ([]dto.NewLogResponseAll, *errs.AppError) {
	return s.repo.FindAll(id)
}
func (s DefaultLogService) GetLogsByDate(id string, date string) ([]dto.NewLogResponseByDate, *errs.AppError) {
	return s.repo.FindByDate(id, date)
}
func (s DefaultLogService) GetLogsByVisitorId(nic string) ([]dto.NewLogResponseByNic, *errs.AppError) {
	return s.repo.FindByVisitorId(nic)
}
func NewLogService(repo domain.LogRepository) DefaultLogService {
	return DefaultLogService{repo}
}
