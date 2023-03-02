package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"school-marks-app/helpers"
	"school-marks-app/internal/app/mapper"
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/internal/app/model/response"
	"school-marks-app/internal/app/repository"
	"school-marks-app/pkg/errorHandler"
)

type Student interface {
	GetById(id uint) (*response.Student, *errorHandler.WebError)
	Create(newStudent *entity.Student) (*response.Student, *errorHandler.WebError)
	Update(newStudent *entity.Student) (*response.Student, *errorHandler.WebError)
	Delete(id uint) *errorHandler.WebError
	BulkCreate(students []entity.Student) ([]uint, *errorHandler.WebError)
}

type StudentService struct {
	studentRepo  repository.Student
	classService Class
}

func NewStudentService(studentRepo repository.Student, classService Class) *StudentService {
	return &StudentService{studentRepo: studentRepo, classService: classService}
}

func (s *StudentService) ValidateClassExists(student *entity.Student) *errorHandler.WebError {
	if _, webErr := s.classService.GetById(student.ClassID); webErr != nil {
		return webErr
	}
	return nil
}

func (s *StudentService) GetById(id uint) (*response.Student, *errorHandler.WebError) {
	student, webErr := s.studentRepo.GetById(id)
	if webErr != nil {
		return nil, webErr
	}

	if student.ID == 0 {
		return nil, &errorHandler.WebError{
			Err:  errors.New(fmt.Sprintf("student with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}

	studentResponse := mapper.StudentToStudentResponse(student)
	return studentResponse, nil
}

func (s *StudentService) Create(newStudent *entity.Student) (*response.Student, *errorHandler.WebError) {
	if webErr := s.ValidateClassExists(newStudent); webErr != nil {
		return nil, webErr
	}
	newStudent, webErr := s.studentRepo.Create(newStudent)
	if webErr != nil {
		return nil, webErr
	}

	studentResponse := mapper.StudentToStudentResponse(newStudent)
	return studentResponse, nil
}

func (s *StudentService) Update(newStudent *entity.Student) (*response.Student, *errorHandler.WebError) {
	oldStudent, webErr := s.GetById(newStudent.ID)
	if webErr != nil {
		return nil, webErr
	}

	if webErr = s.ValidateClassExists(newStudent); webErr != nil {
		return nil, webErr
	}

	newStudent.ParentId = oldStudent.ID
	newStudent.ID = 0

	newStudent, webErr = s.studentRepo.Update(newStudent, oldStudent.ID)
	if webErr != nil {
		return nil, webErr
	}

	studentResponse := mapper.StudentToStudentResponse(newStudent)
	return studentResponse, nil
}

func (s *StudentService) Delete(id uint) *errorHandler.WebError {
	student, webErr := s.GetById(id)
	if webErr != nil {
		return webErr
	}
	s.studentRepo.Delete(student.ID)

	return nil
}

func (s *StudentService) BulkCreate(students []entity.Student) ([]uint, *errorHandler.WebError) {
	validatedClassIds := make(map[uint]bool)
	var newIds []uint

	err := helpers.GetConnectionOrCreateAndGet().Transaction(func(tx *gorm.DB) error {
		for _, student := range students {
			if !validatedClassIds[student.ClassID] {
				if webErr := s.ValidateClassExists(&student); webErr != nil {
					return webErr.Err
				}
				validatedClassIds[student.ClassID] = true
			}
			tx.Save(&student)
			newIds = append(newIds, student.ID)
		}
		return nil
	})
	if err != nil {
		return nil, &errorHandler.WebError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	return newIds, nil
}
