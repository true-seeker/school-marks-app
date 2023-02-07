package db

import (
	"fmt"
	errorHandler "school-marks-app/api/error"
	"school-marks-app/api/models"
)

func Migrate() {
	db := GetDBConnection()
	dbInstance, _ := db.conn.DB()
	defer dbInstance.Close()

	err := db.conn.AutoMigrate(&models.AcademicLevel{}, &models.AcademicYear{}, &models.Class{}, &models.SchoolClass{}, &models.Student{}, &models.Teacher{})
	errorHandler.FailOnError(err, "Failed to migrate")

}

func CreateCatalogs() {
	db := GetDBConnection()
	dbInstance, _ := db.conn.DB()
	defer dbInstance.Close()

	var level models.AcademicLevel
	if err := db.conn.Where("title = ?", "Начальная школа").First(&level).Error; err != nil {
		db.conn.Create(&models.AcademicLevel{
			ID:    0,
			Title: "Начальная школа",
		})
	}
	if err := db.conn.Where("title = ?", "Средняя школа").First(&level).Error; err != nil {
		db.conn.Create(&models.AcademicLevel{
			ID:    1,
			Title: "Средняя школа",
		})
	}

	if err := db.conn.Where("title = ?", "Старшая школа").First(&level).Error; err != nil {
		db.conn.Create(&models.AcademicLevel{
			ID:    2,
			Title: "Старшая школа",
		})
	}

	var class models.SchoolClass
	for i := 0; i < 11; i++ {
		if err := db.conn.Where("title = ?", fmt.Sprintf("%d", i+1)).First(&class).Error; err != nil {
			db.conn.Create(&models.SchoolClass{
				ID:    uint(i),
				Title: fmt.Sprintf("%d", i+1),
			})
		}
	}
}
