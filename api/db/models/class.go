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
	ParentId      uint         `json:"parent_id"`
	TeacherID     uint         `json:"teacher_id"`
	Teacher       Teacher      `json:"teacher"`
	SchoolClassId uint         `json:"school_class_id"`
	SchoolClass   SchoolClass  `json:"school_class"`
	YearID        uint         `json:"year_id"`
	Year          AcademicYear `json:"year"`
}

func ValidateClassExistingEntities(class Class) *error2.WebError {
	var teacherModel Teacher
	var yearModel AcademicYear
	var schoolClassModel SchoolClass

	if _, webErr := teacherModel.GetById(class.TeacherID); webErr != nil {
		return webErr
	}

	if _, webErr := yearModel.GetById(class.YearID); webErr != nil {
		return webErr
	}

	if _, webErr := schoolClassModel.GetById(class.SchoolClassId); webErr != nil {
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

	if webErr := ValidateClassExistingEntities(c); webErr != nil {
		return nil, webErr
	}

	dbConnection.Create(&c)
	dbConnection.Preload(clause.Associations).Find(&c)
	return &c, nil
}

func (c Class) Update() (*Class, *error2.WebError) {
	dbConnection := db.GetDB()
	oldClass, webErr := c.GetById(c.ID)
	if webErr != nil {
		return nil, webErr
	}

	if webErr = ValidateClassExistingEntities(c); webErr != nil {
		return nil, webErr
	}

	c.ParentId = oldClass.ID
	c.ID = 0
	dbConnection.Delete(&oldClass, oldClass.ID)
	dbConnection.Create(&c)
	dbConnection.Model(&Student{}).Where("class_id = ?", c.ParentId).Update("class_id", c.ID)

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
