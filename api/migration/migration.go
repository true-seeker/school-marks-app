package migration

import (
	"fmt"
	db3 "school-marks-app/api/db"
	db2 "school-marks-app/api/db/models"
	errorHandler "school-marks-app/api/error"
)

func Migrate() {
	dbConnection := db3.GetDB()

	err := dbConnection.AutoMigrate(&db2.AcademicLevel{}, &db2.AcademicYear{}, &db2.Class{}, &db2.SchoolClass{}, &db2.Student{}, &db2.Teacher{})
	errorHandler.FailOnError(err, "Failed to migrate")

}

func CreateCatalogs() {
	dbConnection := db3.GetDB()

	dbConnection.Save(&db2.AcademicLevel{
		ID:    1,
		Title: "Начальная школа",
	})
	dbConnection.Save(&db2.AcademicLevel{
		ID:    2,
		Title: "Средняя школа",
	})

	dbConnection.Save(&db2.AcademicLevel{
		ID:    3,
		Title: "Старшая школа",
	})

	for i := 1; i <= 11; i++ {
		var levelId uint
		if i <= 4 {
			levelId = 1
		} else if i <= 9 {
			levelId = 2
		} else {
			levelId = 3
		}
		dbConnection.Save(&db2.SchoolClass{
			ID:      uint(i),
			Title:   fmt.Sprintf("%d", i),
			LevelId: levelId,
		})
	}
}
