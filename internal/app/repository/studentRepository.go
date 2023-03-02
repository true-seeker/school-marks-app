package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/pkg/errorHandler"
)

type Student interface {
	GetById(id uint) (*entity.Student, *errorHandler.WebError)
	Create(newStudent *entity.Student) (*entity.Student, *errorHandler.WebError)
	Update(newStudent *entity.Student, oldStudentId uint) (*entity.Student, *errorHandler.WebError)
	Delete(id uint) *errorHandler.WebError
	BulkCreate(students []entity.Student) ([]uint, *errorHandler.WebError)
}

func NewStudentRepository(db *gorm.DB) Student {
	return &StudentRepository{Db: db}
}

type StudentRepository struct {
	Db *gorm.DB
}

func (s *StudentRepository) GetById(id uint) (*entity.Student, *errorHandler.WebError) {
	student := &entity.Student{}
	s.Db.Preload("Class.Teacher").Preload("Class.Year").Preload("Class.SchoolClass").Preload(clause.Associations).First(&student, id)
	return student, nil
}

func (s *StudentRepository) Create(newStudent *entity.Student) (*entity.Student, *errorHandler.WebError) {
	s.Db.Create(&newStudent)
	s.Db.Preload("Class.Teacher").Preload("Class.Year").Preload("Class.SchoolClass").Preload(clause.Associations).Find(&newStudent)
	return newStudent, nil
}

func (s *StudentRepository) Update(newStudent *entity.Student, oldStudentId uint) (*entity.Student, *errorHandler.WebError) {
	s.Db.Delete(&entity.Student{}, oldStudentId)
	s.Db.Create(&newStudent)
	s.Db.Preload("Class.Teacher").Preload("Class.Year").Preload("Class.SchoolClass").Preload(clause.Associations).Find(&newStudent)
	return newStudent, nil
}

func (s *StudentRepository) Delete(id uint) *errorHandler.WebError {
	s.Db.Delete(&entity.Student{}, id)
	return nil
}

func (s *StudentRepository) BulkCreate(students []entity.Student) ([]uint, *errorHandler.WebError) {
	//TODO implement me
	panic("implement me")
}
