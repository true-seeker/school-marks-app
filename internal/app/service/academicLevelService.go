package service

import (
	"school-marks-app/internal/app/mapper"
	"school-marks-app/internal/app/model/response"
	"school-marks-app/internal/app/repository"
	"school-marks-app/pkg/errorHandler"
)

type AcademicLevel interface {
	Get() (*[]response.AcademicLevel, *errorHandler.WebError)
	GetById(id uint) (*response.AcademicLevel, *errorHandler.WebError)
}

type AcademicLevelService struct {
	academicLevelRepo repository.AcademicLevel
}

func NewAcademicLevelService(academicLevelRepo repository.AcademicLevel) AcademicLevel {
	return &AcademicLevelService{academicLevelRepo: academicLevelRepo}
}

func (a *AcademicLevelService) Get() (*[]response.AcademicLevel, *errorHandler.WebError) {
	academicLevel, webErr := a.academicLevelRepo.Get()
	if webErr != nil {
		return nil, webErr
	}
	academicLevelResponse := mapper.AcademicLevelsToAcademicLevelResponses(academicLevel)
	return academicLevelResponse, nil
}

func (a *AcademicLevelService) GetById(id uint) (*response.AcademicLevel, *errorHandler.WebError) {
	academicLevel, webErr := a.academicLevelRepo.GetById(id)
	if webErr != nil {
		return nil, webErr
	}
	academicLevelResponse := mapper.AcademicLevelToAcademicLevelResponse(academicLevel)
	return academicLevelResponse, nil
}
