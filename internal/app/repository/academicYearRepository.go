package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/pkg/errorHandler"
)

type AcademicYear interface {
	GetById(id uint) (*entity.AcademicYear, *errorHandler.WebError)
	GetByYear(academicYear *entity.AcademicYear) *entity.AcademicYear
	Create(newAcademicYear *entity.AcademicYear) (*entity.AcademicYear, *errorHandler.WebError)
	Update(newAcademicYear *entity.AcademicYear, oldAcademicYearId uint) (*entity.AcademicYear, *errorHandler.WebError)
	Delete(id uint) *errorHandler.WebError
}

type AcademicYearRepository struct {
	Db *gorm.DB
}

func NewAcademicYearRepository(db *gorm.DB) AcademicYear {
	return &AcademicYearRepository{Db: db}
}

func (a *AcademicYearRepository) GetById(id uint) (*entity.AcademicYear, *errorHandler.WebError) {
	academicYear := &entity.AcademicYear{}

	a.Db.First(academicYear, id)
	if academicYear.ID == 0 {
		return nil, &errorHandler.WebError{
			Err:  errors.New(fmt.Sprintf("academicYear with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}
	return academicYear, nil
}

func (a *AcademicYearRepository) GetByYear(academicYear *entity.AcademicYear) *entity.AcademicYear {
	a.Db.Where("year = ?", academicYear.Year).First(&academicYear)
	return academicYear
}

func (a *AcademicYearRepository) Create(newAcademicYear *entity.AcademicYear) (*entity.AcademicYear, *errorHandler.WebError) {
	a.Db.Create(newAcademicYear)
	return newAcademicYear, nil
}

func (a *AcademicYearRepository) Update(newAcademicYear *entity.AcademicYear, oldAcademicYearId uint) (*entity.AcademicYear, *errorHandler.WebError) {
	a.Db.Delete(oldAcademicYearId)
	a.Db.Create(newAcademicYear)
	a.Db.Model(&entity.Class{}).Where("year_id = ?", newAcademicYear.ParentId).Update("year_id", newAcademicYear.ID)
	return newAcademicYear, nil
}

func (a *AcademicYearRepository) Delete(id uint) *errorHandler.WebError {
	a.Db.Delete(&entity.AcademicYear{}, id)
	return nil
}
