package migration

import (
	"fmt"
	db3 "school-marks-app/api/db"
	db2 "school-marks-app/api/db/models"
	errorHandler "school-marks-app/api/error"
)

func Migrate() {
	db := db3.GetDBConnection()
	dbInstance, _ := db.DB()
	defer dbInstance.Close()

	err := db.AutoMigrate(&db2.AcademicLevel{}, &db2.AcademicYear{}, &db2.Class{}, &db2.SchoolClass{}, &db2.Student{}, &db2.Teacher{})
	errorHandler.FailOnError(err, "Failed to migrate")

}

func CreateCatalogs() {
	db := db3.GetDBConnection()
	dbInstance, _ := db.DB()
	defer dbInstance.Close()

	var level db2.AcademicLevel
	if err := db.Where("title = ?", "Начальная школа").First(&level).Error; err != nil {
		db.Create(&db2.AcademicLevel{
			ID:    0,
			Title: "Начальная школа",
		})
	}
	if err := db.Where("title = ?", "Средняя школа").First(&level).Error; err != nil {
		db.Create(&db2.AcademicLevel{
			ID:    1,
			Title: "Средняя школа",
		})
	}

	if err := db.Where("title = ?", "Старшая школа").First(&level).Error; err != nil {
		db.Create(&db2.AcademicLevel{
			ID:    2,
			Title: "Старшая школа",
		})
	}

	var class db2.SchoolClass
	for i := 0; i < 11; i++ {
		if err := db.Where("title = ?", fmt.Sprintf("%d", i+1)).First(&class).Error; err != nil {
			db.Create(&db2.SchoolClass{
				ID:    uint(i),
				Title: fmt.Sprintf("%d", i+1),
			})
		}
	}
}
