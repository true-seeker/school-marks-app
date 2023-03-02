package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/pkg/errorHandler"
)

type SchoolClass interface {
	Get() (*[]entity.SchoolClass, *errorHandler.WebError)
	GetById(id uint) (*entity.SchoolClass, *errorHandler.WebError)
}

type SchoolClassRepository struct {
	Db *gorm.DB
}

func NewSchoolClassRepository(db *gorm.DB) SchoolClass {
	return &SchoolClassRepository{Db: db}
}

func (s *SchoolClassRepository) Get() (*[]entity.SchoolClass, *errorHandler.WebError) {
	var schoolClasses *[]entity.SchoolClass
	s.Db.Preload(clause.Associations).Find(&schoolClasses)

	return schoolClasses, nil
}

func (s *SchoolClassRepository) GetById(id uint) (*entity.SchoolClass, *errorHandler.WebError) {
	schoolClass := &entity.SchoolClass{}
	s.Db.First(&schoolClass, id)
	return schoolClass, nil
}
