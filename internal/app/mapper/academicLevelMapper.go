package mapper

import (
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/internal/app/model/response"
)

func AcademicLevelToAcademicLevelResponse(academicLevel *entity.AcademicLevel) *response.AcademicLevel {
	r := &response.AcademicLevel{
		Id:    academicLevel.ID,
		Title: academicLevel.Title,
	}
	return r
}

func AcademicLevelsToAcademicLevelResponses(academicLevels *[]entity.AcademicLevel) *[]response.AcademicLevel {
	rs := make([]response.AcademicLevel, 0)

	for _, academicLevel := range *academicLevels {
		r := AcademicLevelToAcademicLevelResponse(&academicLevel)
		rs = append(rs, *r)
	}

	return &rs
}
