package db

import "gorm.io/gorm"

// Student Ученик
type Student struct {
	gorm.Model
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	ClassID    int    `json:"class_id"`
	Class      Class  `json:"class"`
}
