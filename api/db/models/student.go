package db

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"school-marks-app/api/db"
	error2 "school-marks-app/api/error"
)

// Student Ученик
type Student struct {
	gorm.Model
	ParentId   uint   `json:"parent_id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	ClassID    uint   `json:"class_id"`
	Class      Class  `json:"class"`
}

func ValidateStudentExistingEntities(student Student) *error2.WebError {
	var classModel Class

	if _, webErr := classModel.GetById(student.ClassID); webErr != nil {
		return webErr
	}
	return nil
}

func (s Student) GetById(id uint) (*Student, *error2.WebError) {
	dbConnection := db.GetDB()

	student := Student{}

	dbConnection.First(&student, id)
	if student.ID == 0 {
		return nil, &error2.WebError{
			Err:  errors.New(fmt.Sprintf("student with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}

	return &student, nil
}

func (s Student) Create() (*Student, *error2.WebError) {
	dbConnection := db.GetDB()

	if webErr := ValidateStudentExistingEntities(s); webErr != nil {
		return nil, webErr
	}

	dbConnection.Create(&s)
	dbConnection.Preload("Class.Teacher").Preload("Class.Year").Preload(clause.Associations).Find(&s)

	return &s, nil
}

func (s Student) Update() (*Student, *error2.WebError) {
	dbConnection := db.GetDB()

	oldTeacher, webErr := s.GetById(s.ID)
	if webErr != nil {
		return nil, webErr
	}

	if webErr = ValidateStudentExistingEntities(s); webErr != nil {
		return nil, webErr
	}

	s.ParentId = oldTeacher.ID
	s.ID = 0

	dbConnection.Delete(&oldTeacher, oldTeacher.ID)
	dbConnection.Create(&s)
	dbConnection.Preload("Class.Teacher").Preload("Class.Year").Preload(clause.Associations).Find(&s)
	return &s, nil
}

func (s Student) Delete(id uint) *error2.WebError {
	dbConnection := db.GetDB()

	teacher, webErr := s.GetById(id)
	if webErr != nil {
		return webErr
	}

	dbConnection.Delete(&teacher, id)

	return nil
}
