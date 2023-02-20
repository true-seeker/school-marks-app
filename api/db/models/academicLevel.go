package db

import (
	"errors"
	"fmt"
	"net/http"
	"school-marks-app/api/db"
	error2 "school-marks-app/api/error"
)

// AcademicLevel Академический уровень (нач, ср, старш школа)
type AcademicLevel struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `json:"title"`
}

func (a AcademicLevel) Get() (*[]AcademicLevel, *error2.WebError) {
	dbConnection := db.GetDB()

	var academicLevels []AcademicLevel

	dbConnection.Find(&academicLevels)

	return &academicLevels, nil
}

func (a AcademicLevel) GetById(id uint) (*AcademicLevel, *error2.WebError) {
	dbConnection := db.GetDB()

	academicLevel := AcademicLevel{}

	dbConnection.First(&academicLevel, id)
	if academicLevel.ID == 0 {
		return nil, &error2.WebError{
			Err:  errors.New(fmt.Sprintf("academicLevel with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}

	return &academicLevel, nil
}
