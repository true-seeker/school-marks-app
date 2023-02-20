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
	ParentId   uint   `json:"parent_id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

func (t Teacher) GetById(id uint) (*Teacher, *error2.WebError) {
	dbConnection := db.GetDB()

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

func (t Teacher) Create() (*Teacher, *error2.WebError) {
	dbConnection := db.GetDB()

	dbConnection.Create(&t)
	return &t, nil
}

func (t Teacher) Update() (*Teacher, *error2.WebError) {
	dbConnection := db.GetDB()

	oldTeacher, webErr := t.GetById(t.ID)
	if webErr != nil {
		return nil, webErr
	}

	t.ParentId = oldTeacher.ID
	t.ID = 0

	dbConnection.Delete(&oldTeacher, oldTeacher.ID)
	dbConnection.Create(&t)
	dbConnection.Model(&Class{}).Where("teacher_id = ?", t.ParentId).Update("teacher_id", t.ID)
	return &t, nil
}

func (t Teacher) Delete(id uint) *error2.WebError {
	dbConnection := db.GetDB()

	teacher, webErr := t.GetById(id)
	if webErr != nil {
		return webErr
	}

	dbConnection.Delete(&teacher, id)

	return nil
}
