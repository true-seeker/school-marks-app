package models

import "gorm.io/gorm"

// Teacher Учитель
type Teacher struct {
	gorm.Model
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
