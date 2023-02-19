package db

import "gorm.io/gorm"

// Class Класс
type Class struct {
	gorm.Model
	TeacherID int           `json:"teacher_id"`
	Teacher   Teacher       `json:"teacher"`
	LevelID   int           `json:"level_id"`
	Level     AcademicLevel `json:"level"`
	YearID    int           `json:"year_id"`
	Year      AcademicYear  `json:"year"`
}
