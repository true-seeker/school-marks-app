package response

type Class struct {
	Id          uint         `json:"id"`
	ParentId    uint         `json:"parent_id"`
	Teacher     Teacher      `json:"teacher"`
	SchoolClass SchoolClass  `json:"school_class"`
	Year        AcademicYear `json:"year"`
	Letter      string       `json:"letter"`
}
