package entity

// SchoolClass Параллель
type SchoolClass struct {
	ID      uint          `gorm:"primaryKey"`
	Title   string        `json:"title"`
	LevelId uint          `json:"level_id"`
	Level   AcademicLevel `json:"level"`
}
