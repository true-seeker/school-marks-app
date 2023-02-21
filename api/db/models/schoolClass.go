package db

import (
	"errors"
	"fmt"
	"net/http"
	"school-marks-app/api/db"
	error2 "school-marks-app/api/error"
)

// SchoolClass Параллель
type SchoolClass struct {
	ID      uint          `gorm:"primaryKey"`
	Title   string        `json:"title"`
	LevelId uint          `json:"level_id"`
	Level   AcademicLevel `json:"level"`
}

func (s SchoolClass) Get() (*[]SchoolClass, *error2.WebError) {
	dbConnection := db.GetDB()

	var schoolClasses []SchoolClass

	dbConnection.Find(&schoolClasses)

	return &schoolClasses, nil
}

func (s SchoolClass) GetById(id uint) (*SchoolClass, *error2.WebError) {
	dbConnection := db.GetDB()

	schoolClass := SchoolClass{}

	dbConnection.First(&schoolClass, id)
	if schoolClass.ID == 0 {
		return nil, &error2.WebError{
			Err:  errors.New(fmt.Sprintf("schoolClass with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}

	return &schoolClass, nil
}
