package db

import (
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
