package mapper

import (
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/internal/app/model/response"
)

func SchoolClassToSchoolClassResponse(schoolClass *entity.SchoolClass) *response.SchoolClass {
	r := &response.SchoolClass{
		ID:    schoolClass.ID,
		Title: schoolClass.Title,
		Level: *AcademicLevelToAcademicLevelResponse(&schoolClass.Level),
	}
	return r
}

func SchoolClassesToSchoolClassResponses(schoolClasses *[]entity.SchoolClass) *[]response.SchoolClass {
	var rs []response.SchoolClass
	for _, schoolClass := range *schoolClasses {
		r := SchoolClassToSchoolClassResponse(&schoolClass)
		rs = append(rs, *r)
	}
	return &rs
}
