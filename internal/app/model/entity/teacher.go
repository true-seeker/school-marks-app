package entity

import (
	"gorm.io/gorm"
)

// Teacher Учитель
type Teacher struct {
	gorm.Model
	ParentId   uint   `json:"parent_id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}
