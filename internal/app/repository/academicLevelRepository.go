package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/pkg/errorHandler"
)

type AcademicLevel interface {
	Get() (*[]entity.AcademicLevel, *errorHandler.WebError)
	GetById(id uint) (*entity.AcademicLevel, *errorHandler.WebError)
}

type AcademicLevelRepository struct {
	Db *gorm.DB
}

func (a AcademicLevelRepository) Get() (*[]entity.AcademicLevel, *errorHandler.WebError) {
	academicLevels := &[]entity.AcademicLevel{}

	a.Db.Find(academicLevels)

	return academicLevels, nil
}

func (a AcademicLevelRepository) GetById(id uint) (*entity.AcademicLevel, *errorHandler.WebError) {
	academicLevel := &entity.AcademicLevel{}

	a.Db.First(academicLevel, id)
	if academicLevel.ID == 0 {
		return nil, &errorHandler.WebError{
			Err:  errors.New(fmt.Sprintf("academicLevel with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}
	return academicLevel, nil
}

func NewAcademicLevelRepository(db *gorm.DB) AcademicLevel {
	return &AcademicLevelRepository{Db: db}
}
