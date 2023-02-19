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
	Year string `json:"year"`
}

func (a AcademicYear) Get(id uint) (*AcademicYear, *error2.WebError) {
	dbConnection := db.GetDBConnection()
	dbInstance, _ := dbConnection.DB()
	defer dbInstance.Close()

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
	dbConnection := db.GetDBConnection()
	dbInstance, _ := dbConnection.DB()
	defer dbInstance.Close()

	var academicYear AcademicYear

	dbConnection.Where("year = ?", a.Year).First(&academicYear)

	if academicYear.ID == 0 {
		return false
	}

	return true
}

func (a AcademicYear) isExistsWithId() bool {
	dbConnection := db.GetDBConnection()
	dbInstance, _ := dbConnection.DB()
	defer dbInstance.Close()

	var academicYear AcademicYear

	dbConnection.First(&academicYear, a.ID)

	if academicYear.ID == 0 {
		return false
	}

	return true
}

func (a AcademicYear) Create() (*AcademicYear, *error2.WebError) {
	dbConnection := db.GetDBConnection()
	dbInstance, _ := dbConnection.DB()
	defer dbInstance.Close()
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
	dbConnection := db.GetDBConnection()
	dbInstance, _ := dbConnection.DB()
	defer dbInstance.Close()

	oldAcademicYear, webErr := a.Get(a.ID)
	if webErr != nil {
		return nil, webErr
	}

	a.CreatedAt = oldAcademicYear.CreatedAt
	dbConnection.Save(&a)
	return &a, nil
}

func (a AcademicYear) Delete(id uint) *error2.WebError {
	dbConnection := db.GetDBConnection()
	dbInstance, _ := dbConnection.DB()
	defer dbInstance.Close()

	academicYear, webErr := a.Get(id)
	if webErr != nil {
		return webErr
	}

	dbConnection.Delete(&academicYear, id)

	return nil
}
