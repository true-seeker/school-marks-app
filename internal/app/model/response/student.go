package response

type Student struct {
	ID         uint   `gorm:"primaryKey"`
	ParentId   uint   `json:"parent_id"`
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
	Class      Class  `json:"class"`
}
