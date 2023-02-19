package db

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"school-marks-app/api/db"
	error2 "school-marks-app/api/error"
)

// Teacher Учитель
type Teacher struct {
	gorm.Model
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

func (t Teacher) Get(id uint) (*Teacher, *error2.WebError) {
	dbConnection := db.GetDBConnection()
	dbInstance, _ := dbConnection.DB()
	defer dbInstance.Close()

	teacher := Teacher{}

	dbConnection.First(&teacher, id)
	if teacher.ID == 0 {
		return nil, &error2.WebError{
			Err:  errors.New(fmt.Sprintf("teacher with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}

	return &teacher, nil
}

func (t Teacher) Create(teacher Teacher) (*Teacher, *error2.WebError) {
	dbConnection := db.GetDBConnection()
	dbInstance, _ := dbConnection.DB()
	defer dbInstance.Close()

	dbConnection.Create(&teacher)
	return &teacher, nil
}

func (t Teacher) Update(teacher Teacher) (*Teacher, *error2.WebError) {
	dbConnection := db.GetDBConnection()
	dbInstance, _ := dbConnection.DB()
	defer dbInstance.Close()

	oldTeacher := Teacher{}

	dbConnection.First(&oldTeacher, teacher.ID)
	if oldTeacher.ID == 0 {
		return nil, &error2.WebError{
			Err:  errors.New(fmt.Sprintf("teacher with id %d does not exist", teacher.ID)),
			Code: http.StatusNotFound,
		}
	}

	teacher.CreatedAt = oldTeacher.CreatedAt
	dbConnection.Save(&teacher)
	return &teacher, nil
}

func (t Teacher) Delete(id uint) *error2.WebError {
	dbConnection := db.GetDBConnection()
	dbInstance, _ := dbConnection.DB()
	defer dbInstance.Close()

	teacher := Teacher{}

	dbConnection.First(&teacher, id)
	if teacher.ID == 0 {
		return &error2.WebError{
			Err:  errors.New(fmt.Sprintf("teacher with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}

	dbConnection.Delete(&teacher, id)

	return nil
}
