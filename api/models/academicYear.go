package models

import "gorm.io/gorm"

// AcademicYear Академический год
type AcademicYear struct {
	gorm.Model
	Year string `json:"year"`
}
