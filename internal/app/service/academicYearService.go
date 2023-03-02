package service

import (
	"errors"
	"fmt"
	"net/http"
	"school-marks-app/internal/app/mapper"
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/internal/app/model/response"
	"school-marks-app/internal/app/repository"
	"school-marks-app/pkg/errorHandler"
)

type AcademicYear interface {
	GetById(id uint) (*response.AcademicYear, *errorHandler.WebError)
	GetByYear(academicYear *entity.AcademicYear) *entity.AcademicYear
	Create(newAcademicYear *entity.AcademicYear) (*response.AcademicYear, *errorHandler.WebError)
	Update(newAcademicYear *entity.AcademicYear) (*response.AcademicYear, *errorHandler.WebError)
	Delete(id uint) *errorHandler.WebError
}

type AcademicYearService struct {
	academicYearRepo repository.AcademicYear
}

func NewAcademicYearService(academicYearRepo repository.AcademicYear) AcademicYear {
	return &AcademicYearService{academicYearRepo: academicYearRepo}
}

func (a *AcademicYearService) GetById(id uint) (*response.AcademicYear, *errorHandler.WebError) {
	academicYear, webErr := a.academicYearRepo.GetById(id)
	if webErr != nil {
		return nil, webErr
	}
	if academicYear.ID == 0 {
		return nil, &errorHandler.WebError{
			Err:  errors.New(fmt.Sprintf("academicYear with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}

	academicYearResponse := mapper.AcademicYearToAcademicYearResponse(academicYear)

	return academicYearResponse, nil
}

func (a *AcademicYearService) GetByYear(academicYear *entity.AcademicYear) *entity.AcademicYear {
	academicYear = a.academicYearRepo.GetByYear(academicYear)

	return academicYear
}

func (a *AcademicYearService) Create(newAcademicYear *entity.AcademicYear) (*response.AcademicYear, *errorHandler.WebError) {
	if a.academicYearRepo.GetByYear(newAcademicYear).ID != 0 {
		return nil, &errorHandler.WebError{
			Err:  errors.New(fmt.Sprintf("academicYear with year %s already exists", newAcademicYear.Year)),
			Code: http.StatusNotFound,
		}
	}

	newAcademicYear, webErr := a.academicYearRepo.Create(newAcademicYear)
	if webErr != nil {
		return nil, webErr
	}

	academicYearResponse := mapper.AcademicYearToAcademicYearResponse(newAcademicYear)
	return academicYearResponse, nil
}

func (a *AcademicYearService) Update(newAcademicYear *entity.AcademicYear) (*response.AcademicYear, *errorHandler.WebError) {
	oldAcademicYear, webErr := a.GetById(newAcademicYear.ID)
	if webErr != nil {
		return nil, webErr
	}
	newAcademicYear.ParentId = oldAcademicYear.Id
	newAcademicYear.ID = 0
	newAcademicYear, webErr = a.academicYearRepo.Update(newAcademicYear, oldAcademicYear.Id)
	if webErr != nil {
		return nil, webErr
	}

	academicYearResponse := mapper.AcademicYearToAcademicYearResponse(newAcademicYear)
	return academicYearResponse, nil
}

func (a *AcademicYearService) Delete(id uint) *errorHandler.WebError {
	academicYear, webErr := a.GetById(id)
	if webErr != nil {
		return webErr
	}
	a.academicYearRepo.Delete(academicYear.Id)

	return nil
}
