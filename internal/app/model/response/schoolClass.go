package response

type SchoolClass struct {
	ID    uint          `gorm:"primaryKey"`
	Title string        `json:"title"`
	Level AcademicLevel `json:"level"`
}
