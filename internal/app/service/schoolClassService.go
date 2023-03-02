package service

import (
	"errors"
	"fmt"
	"net/http"
	"school-marks-app/internal/app/mapper"
	"school-marks-app/internal/app/model/response"
	"school-marks-app/internal/app/repository"
	error2 "school-marks-app/pkg/errorHandler"
)

type SchoolClass interface {
	Get() (*[]response.SchoolClass, *error2.WebError)
	GetById(id uint) (*response.SchoolClass, *error2.WebError)
}

type SchoolClassService struct {
	schoolClassRepo repository.SchoolClass
}

func NewSchoolClassService(schoolClassRepo repository.SchoolClass) SchoolClass {
	return &SchoolClassService{schoolClassRepo: schoolClassRepo}
}

func (s *SchoolClassService) Get() (*[]response.SchoolClass, *error2.WebError) {
	schoolClasses, webErr := s.schoolClassRepo.Get()
	if webErr != nil {
		return nil, webErr
	}
	schoolClassResponses := mapper.SchoolClassesToSchoolClassResponses(schoolClasses)
	return schoolClassResponses, nil
}

func (s *SchoolClassService) GetById(id uint) (*response.SchoolClass, *error2.WebError) {
	schoolClass, webErr := s.schoolClassRepo.GetById(id)
	if webErr != nil {
		return nil, webErr
	}

	if schoolClass.ID == 0 {
		return nil, &error2.WebError{
			Err:  errors.New(fmt.Sprintf("schoolClass with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}
	schoolClassResponse := mapper.SchoolClassToSchoolClassResponse(schoolClass)

	return schoolClassResponse, nil
}
