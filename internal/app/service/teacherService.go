package service

import (
	"errors"
	"fmt"
	"net/http"
	"school-marks-app/internal/app/mapper"
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/internal/app/model/response"
	"school-marks-app/internal/app/repository"
	error2 "school-marks-app/pkg/errorHandler"
)

type Teacher interface {
	GetById(id uint) (*response.Teacher, *error2.WebError)
	Create(teacher *entity.Teacher) (*response.Teacher, *error2.WebError)
	Update(teacher *entity.Teacher) (*response.Teacher, *error2.WebError)
	Delete(id uint) *error2.WebError
}

type TeacherService struct {
	teacherRepo repository.Teacher
}

func NewTeacherService(teacherRepo repository.Teacher) Teacher {
	return &TeacherService{teacherRepo: teacherRepo}
}

func (t *TeacherService) GetById(id uint) (*response.Teacher, *error2.WebError) {
	teacher, webErr := t.teacherRepo.GetById(id)
	if webErr != nil {
		return nil, webErr
	}
	if teacher.ID == 0 {
		return nil, &error2.WebError{
			Err:  errors.New(fmt.Sprintf("teacher with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}
	teacherResponse := mapper.TeacherToTeacherResponse(teacher)
	return teacherResponse, nil
}

func (t *TeacherService) Create(teacher *entity.Teacher) (*response.Teacher, *error2.WebError) {
	newTeacher, webErr := t.teacherRepo.Create(teacher)
	if webErr != nil {
		return nil, webErr
	}
	teacherResponse := mapper.TeacherToTeacherResponse(newTeacher)
	return teacherResponse, nil
}

func (t *TeacherService) Update(newTeacher *entity.Teacher) (*response.Teacher, *error2.WebError) {
	oldTeacher, webErr := t.GetById(newTeacher.ID)
	if webErr != nil {
		return nil, webErr
	}

	newTeacher.ParentId = oldTeacher.ID
	newTeacher.ID = 0

	teacher, webErr := t.teacherRepo.Update(newTeacher, oldTeacher.ID)
	if webErr != nil {
		return nil, webErr
	}

	teacherResponse := mapper.TeacherToTeacherResponse(teacher)
	return teacherResponse, nil
}

func (t *TeacherService) Delete(id uint) *error2.WebError {
	_, webErr := t.GetById(id)
	if webErr != nil {
		return webErr
	}
	t.teacherRepo.Delete(id)
	return nil
}
