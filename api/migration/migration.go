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

	var level db2.AcademicLevel
	if err := dbConnection.Where("title = ?", "Начальная школа").First(&level).Error; err != nil {
		dbConnection.Create(&db2.AcademicLevel{
			ID:    1,
			Title: "Начальная школа",
		})
	}
	if err := dbConnection.Where("title = ?", "Средняя школа").First(&level).Error; err != nil {
		dbConnection.Create(&db2.AcademicLevel{
			ID:    2,
			Title: "Средняя школа",
		})
	}

	if err := dbConnection.Where("title = ?", "Старшая школа").First(&level).Error; err != nil {
		dbConnection.Create(&db2.AcademicLevel{
			ID:    3,
			Title: "Старшая школа",
		})
	}

	var class db2.SchoolClass
	for i := 0; i < 11; i++ {
		if err := dbConnection.Where("title = ?", fmt.Sprintf("%d", i+1)).First(&class).Error; err != nil {
			dbConnection.Create(&db2.SchoolClass{
				ID:    uint(i),
				Title: fmt.Sprintf("%d", i+1),
			})
		}
	}
}
