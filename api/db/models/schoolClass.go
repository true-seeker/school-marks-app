package db

import (
	"school-marks-app/api/db"
	error2 "school-marks-app/api/error"
)

// SchoolClass Параллель
type SchoolClass struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `json:"title"`
}

func (a SchoolClass) Get() (*[]SchoolClass, *error2.WebError) {
	dbConnection := db.GetDBConnection()
	dbInstance, _ := dbConnection.DB()
	defer dbInstance.Close()

	var schoolClasses []SchoolClass

	dbConnection.Find(&schoolClasses)

	return &schoolClasses, nil
}
