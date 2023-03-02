package migration

import (
	"fmt"
	"school-marks-app/helpers"
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/pkg/errorHandler"
)

func Migrate() {
	dbConnection := helpers.GetConnectionOrCreateAndGet()

	err := dbConnection.AutoMigrate(&entity.AcademicLevel{}, &entity.AcademicYear{}, &entity.Class{}, &entity.SchoolClass{}, &entity.Student{}, &entity.Teacher{})
	errorHandler.FailOnError(err, "Failed to migrate")

}

func CreateCatalogs() {
	dbConnection := helpers.GetConnectionOrCreateAndGet()

	dbConnection.Save(&entity.AcademicLevel{
		ID:    1,
		Title: "Начальная школа",
	})
	dbConnection.Save(&entity.AcademicLevel{
		ID:    2,
		Title: "Средняя школа",
	})

	dbConnection.Save(&entity.AcademicLevel{
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
		dbConnection.Save(&entity.SchoolClass{
			ID:      uint(i),
			Title:   fmt.Sprintf("%d", i),
			LevelId: levelId,
		})
	}
}
