package mapper

import (
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/internal/app/model/response"
)

func TeacherToTeacherResponse(teacher *entity.Teacher) *response.Teacher {
	r := &response.Teacher{
		ID:         teacher.ID,
		ParentId:   teacher.ParentId,
		Name:       teacher.Name,
		Surname:    teacher.Surname,
		Patronymic: teacher.Patronymic,
	}
	return r
}

func TeachersToTeacherResponses(teachers *[]entity.Teacher) *[]response.Teacher {
	rs := make([]response.Teacher, 0)
	for _, teacher := range *teachers {
		r := TeacherToTeacherResponse(&teacher)
		rs = append(rs, *r)
	}
	return &rs
}
