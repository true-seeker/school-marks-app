package mapper

import (
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/internal/app/model/response"
)

func ClassToClassResponse(class *entity.Class) *response.Class {
	r := &response.Class{
		Id:          class.ID,
		ParentId:    class.ParentId,
		Teacher:     *TeacherToTeacherResponse(&class.Teacher),
		SchoolClass: *SchoolClassToSchoolClassResponse(&class.SchoolClass),
		Year:        *AcademicYearToAcademicYearResponse(&class.Year),
		Letter:      class.Letter,
	}

	return r
}
