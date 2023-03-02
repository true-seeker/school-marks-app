package entity

import (
	"gorm.io/gorm"
)

// AcademicYear Академический год
type AcademicYear struct {
	gorm.Model
	ParentId uint   `json:"parent_id"`
	Year     string `json:"year"`
}
