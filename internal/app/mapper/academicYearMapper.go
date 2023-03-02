package mapper

import (
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/internal/app/model/response"
)

func AcademicYearToAcademicYearResponse(academicYear *entity.AcademicYear) *response.AcademicYear {
	r := &response.AcademicYear{
		Id:       academicYear.ID,
		ParentId: academicYear.ParentId,
		Year:     academicYear.Year,
	}
	return r
}

func AcademicYearsToAcademicYearResponses(academicLevels *[]entity.AcademicYear) *[]response.AcademicYear {
	rs := make([]response.AcademicYear, 0)

	for _, academicYear := range *academicLevels {
		r := &response.AcademicYear{
			Id:       academicYear.ID,
			ParentId: academicYear.ParentId,
			Year:     academicYear.Year,
		}
		rs = append(rs, *r)
	}

	return &rs
}
