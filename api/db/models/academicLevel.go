package db

// AcademicLevel Академический уровень (нач, ср, старш школа)
type AcademicLevel struct {
	ID    uint   `gorm:"primaryKey"`
	Title string `json:"title"`
}
