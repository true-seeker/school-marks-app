package db

import (
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

func (a SchoolClass) Get() (*[]SchoolClass, *error2.WebError) {
	dbConnection := db.GetDB()

	var schoolClasses []SchoolClass

	dbConnection.Find(&schoolClasses)

	return &schoolClasses, nil
}
