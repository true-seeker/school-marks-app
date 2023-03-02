package response

type AcademicYear struct {
	Id       uint   `json:"id"`
	ParentId uint   `json:"parent_id"`
	Year     string `json:"year"`
}
