package db

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"school-marks-app/api/db"
	error2 "school-marks-app/api/error"
)

// AcademicYear Академический год
type AcademicYear struct {
	gorm.Model
	ParentId uint   `json:"parent_id"`
	Year     string `json:"year"`
}

func (a AcademicYear) GetById(id uint) (*AcademicYear, *error2.WebError) {
	dbConnection := db.GetDB()

	academicYear := AcademicYear{}

	dbConnection.First(&academicYear, id)
	if academicYear.ID == 0 {
		return nil, &error2.WebError{
			Err:  errors.New(fmt.Sprintf("academicYear with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}

	return &academicYear, nil
}

func (a AcademicYear) isExistsWithYear() bool {
	dbConnection := db.GetDB()

	var academicYear AcademicYear

	dbConnection.Where("year = ?", a.Year).First(&academicYear)

	if academicYear.ID == 0 {
		return false
	}

	return true
}

func (a AcademicYear) isExistsWithId() bool {
	dbConnection := db.GetDB()

	var academicYear AcademicYear

	dbConnection.First(&academicYear, a.ID)

	if academicYear.ID == 0 {
		return false
	}

	return true
}

func (a AcademicYear) Create() (*AcademicYear, *error2.WebError) {
	dbConnection := db.GetDB()

	if a.isExistsWithYear() {
		return nil, &error2.WebError{
			Err:  errors.New(fmt.Sprintf("academicYear with year %s already exists", a.Year)),
			Code: http.StatusNotFound,
		}
	}

	dbConnection.Create(&a)
	return &a, nil
}

func (a AcademicYear) Update() (*AcademicYear, *error2.WebError) {
	dbConnection := db.GetDB()

	oldAcademicYear, webErr := a.GetById(a.ID)
	if webErr != nil {
		return nil, webErr
	}
	a.ParentId = oldAcademicYear.ID
	a.ID = 0

	dbConnection.Delete(&oldAcademicYear, oldAcademicYear.ID)
	dbConnection.Create(&a)
	dbConnection.Model(&Class{}).Where("year_id = ?", a.ParentId).Update("year_id", a.ID)
	return &a, nil
}

func (a AcademicYear) Delete(id uint) *error2.WebError {
	dbConnection := db.GetDB()

	academicYear, webErr := a.GetById(id)
	if webErr != nil {
		return webErr
	}

	dbConnection.Delete(&academicYear, id)

	return nil
}
