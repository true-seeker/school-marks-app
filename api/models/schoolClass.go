package models

// SchoolClass Параллель
type SchoolClass struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `json:"title"`
}
