package entity

import (
	"gorm.io/gorm"
)

// Class Класс
type Class struct {
	gorm.Model
	ParentId      uint         `json:"parent_id"`
	TeacherID     uint         `json:"teacher_id"`
	Teacher       Teacher      `json:"teacher"`
	SchoolClassId uint         `json:"school_class_id"`
	SchoolClass   SchoolClass  `json:"school_class"`
	YearID        uint         `json:"year_id"`
	Year          AcademicYear `json:"year"`
	Letter        string       `json:"letter"`
}
