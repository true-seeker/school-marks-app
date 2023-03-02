package mapper

import (
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/internal/app/model/response"
)

func StudentToStudentResponse(student *entity.Student) *response.Student {
	r := &response.Student{
		ID:         student.ID,
		ParentId:   student.ParentId,
		Name:       student.Name,
		Surname:    student.Surname,
		Patronymic: student.Patronymic,
		Class:      *ClassToClassResponse(&student.Class),
	}

	return r
}

func StudentsToStudentResponses(students *[]entity.Student) *[]response.Student {
	rs := make([]response.Student, 0)

	for _, student := range *students {
		r := StudentToStudentResponse(&student)
		rs = append(rs, *r)
	}

	return &rs
}
