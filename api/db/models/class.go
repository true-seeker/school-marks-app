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

// Class Класс
type Class struct {
	gorm.Model
	ParentId  uint          `json:"parent_id"`
	TeacherID uint          `json:"teacher_id"`
	Teacher   Teacher       `json:"teacher"`
	LevelID   uint          `json:"level_id"`
	Level     AcademicLevel `json:"level"`
	YearID    uint          `json:"year_id"`
	Year      AcademicYear  `json:"year"`
}

func ValidateClassExistingEntities(class Class) *error2.WebError {
	var teacherModel Teacher
	var yearModel AcademicYear
	var levelModel AcademicLevel

	_, webErr := teacherModel.GetById(class.TeacherID)
	if webErr != nil {
		return webErr
	}

	_, webErr = yearModel.GetById(class.YearID)
	if webErr != nil {
		return webErr
	}

	_, webErr = levelModel.GetById(class.LevelID)
	if webErr != nil {
		return webErr
	}
	return nil
}

func (c Class) GetById(id uint) (*Class, *error2.WebError) {
	dbConnection := db.GetDB()

	var class Class

	dbConnection.First(&class, id)
	if class.ID == 0 {
		return nil, &error2.WebError{
			Err:  errors.New(fmt.Sprintf("class with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}

	return &class, nil
}

func (c Class) Create() (*Class, *error2.WebError) {
	dbConnection := db.GetDB()

	webErr := ValidateClassExistingEntities(c)
	if webErr != nil {
		return nil, webErr
	}

	dbConnection.Create(&c)
	dbConnection.Preload(clause.Associations).Find(&c)
	return &c, nil
}

func (c Class) Update() (*Class, *error2.WebError) {
	dbConnection := db.GetDB()

	webErr := ValidateClassExistingEntities(c)
	if webErr != nil {
		return nil, webErr
	}

	oldClass, _ := c.GetById(c.ID)

	c.CreatedAt = oldClass.CreatedAt
	dbConnection.Save(&c)
	dbConnection.Preload(clause.Associations).Find(&c)

	return &c, nil
}

func (c Class) Delete(id uint) *error2.WebError {
	dbConnection := db.GetDB()

	class, webErr := c.GetById(id)
	if webErr != nil {
		return webErr
	}

	dbConnection.Delete(&class, id)

	return nil
}
