package service

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"school-marks-app/helpers"
	"school-marks-app/internal/app/mapper"
	"school-marks-app/internal/app/model/entity"
	db2 "school-marks-app/internal/app/model/input"
	"school-marks-app/internal/app/model/response"
	"school-marks-app/internal/app/repository"
	error2 "school-marks-app/pkg/errorHandler"
)

type Class interface {
	GetById(id uint) (*response.Class, *error2.WebError)
	Create(class *entity.Class) (*response.Class, *error2.WebError)
	Update(class *entity.Class) (*response.Class, *error2.WebError)
	Delete(id uint) *error2.WebError
	BulkCreate(classes *[]entity.Class) ([]uint, *error2.WebError)
	Transfer(classTransferInput db2.ClassTransferInput, classId uint) *error2.WebError
}

type ClassService struct {
	classRepo           repository.Class
	studentRepo         repository.Student
	teacherService      Teacher
	academicYearService AcademicYear
	schoolClassService  SchoolClass
}

func NewClassService(classRepo repository.Class, studentRepo repository.Student, teacherService Teacher, academicYearService AcademicYear, schoolClassService SchoolClass) *ClassService {
	return &ClassService{classRepo: classRepo, studentRepo: studentRepo, teacherService: teacherService, academicYearService: academicYearService, schoolClassService: schoolClassService}
}

func (c *ClassService) ValidateClassExistingEntities(class *entity.Class) *error2.WebError {
	if _, webErr := c.teacherService.GetById(class.TeacherID); webErr != nil {
		return webErr
	}

	if _, webErr := c.academicYearService.GetById(class.YearID); webErr != nil {
		return webErr
	}

	if _, webErr := c.schoolClassService.GetById(class.SchoolClassId); webErr != nil {
		return webErr
	}
	return nil
}

func (c *ClassService) GetById(id uint) (*response.Class, *error2.WebError) {
	class, webErr := c.classRepo.GetById(id)
	if webErr != nil {
		return nil, webErr
	}
	classResponse := mapper.ClassToClassResponse(class)
	return classResponse, nil
}

func (c *ClassService) Create(newClass *entity.Class) (*response.Class, *error2.WebError) {
	if webErr := c.ValidateClassExistingEntities(newClass); webErr != nil {
		return nil, webErr
	}
	class, webErr := c.classRepo.Create(newClass)
	if webErr != nil {
		return nil, webErr
	}
	classResponse := mapper.ClassToClassResponse(class)
	return classResponse, nil
}

func (c *ClassService) Update(newClass *entity.Class) (*response.Class, *error2.WebError) {
	if webErr := c.ValidateClassExistingEntities(newClass); webErr != nil {
		return nil, webErr
	}
	oldClass, webErr := c.GetById(newClass.ID)
	newClass.ParentId = oldClass.Id
	newClass.ID = 0

	class, webErr := c.classRepo.Update(newClass, oldClass.Id)
	if webErr != nil {
		return nil, webErr
	}
	classResponse := mapper.ClassToClassResponse(class)
	return classResponse, nil
}

func (c *ClassService) Delete(id uint) *error2.WebError {
	f, webErr := c.GetById(id)
	if webErr != nil {
		return webErr
	}
	if f.Id == 0 {
		return &error2.WebError{
			Err:  errors.New(fmt.Sprintf("class with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}

	webErr = c.classRepo.Delete(id)
	if webErr != nil {
		return webErr
	}

	return nil
}

func (c *ClassService) BulkCreate(classes *[]entity.Class) ([]uint, *error2.WebError) {
	validateTeacherId := make(map[uint]bool)
	validateYearId := make(map[uint]bool)
	validateSchoolClassId := make(map[uint]bool)
	var newIds []uint

	for _, class := range *classes {
		if !validateTeacherId[class.TeacherID] || !validateYearId[class.YearID] || !validateSchoolClassId[class.SchoolClassId] {
			if webErr := c.ValidateClassExistingEntities(&class); webErr != nil {
				return nil, webErr
			}
			validateTeacherId[class.TeacherID] = true
			validateYearId[class.YearID] = true
			validateSchoolClassId[class.SchoolClassId] = true
		}
	}

	newIds, webErr := c.classRepo.BulkCreate(classes)

	if webErr != nil {
		return nil, webErr
	}

	return newIds, nil
}

func (c *ClassService) Transfer(classTransferInput db2.ClassTransferInput, classId uint) *error2.WebError {
	var classStudents []entity.Student
	var classStudentIds, notTransferredStudentIds map[uint]bool
	var db = helpers.GetConnectionOrCreateAndGet()

	db.Where("class_id = ?", classId).Find(&classStudents)
	for _, student := range classStudents {
		classStudentIds[student.ID] = true
	}

	for _, id := range classTransferInput.NotTransferredStudentIds {
		if _, webErr := c.studentRepo.GetById(id); webErr != nil {
			return webErr
		}
		if !classStudentIds[id] {
			return &error2.WebError{
				Err:  errors.New(fmt.Sprintf("Student with id %d does not exist in class", id)),
				Code: http.StatusBadRequest,
			}
		}
	}

	for _, id := range classTransferInput.NewStudentIds {
		if _, webErr := c.studentRepo.GetById(id); webErr != nil {
			return webErr
		}
		notTransferredStudentIds[id] = true
	}

	err := db.Transaction(func(tx *gorm.DB) error {
		newClass := entity.Class{}
		newClass.ID = 0
		newClass.YearID = classTransferInput.NewAcademicYearId
		newClass.SchoolClassId = classTransferInput.NewSchoolClassId
		tx.Create(newClass)

		for _, oldStudent := range classStudents {
			if notTransferredStudentIds[oldStudent.ID] {
				continue
			}

			newStudent := oldStudent
			newStudent.ID = 0
			tx.Delete(&oldStudent)
			newStudent.ClassID = newClass.ID
			tx.Save(newStudent)
		}

		for _, newStudentId := range classTransferInput.NewStudentIds {
			var newStudent entity.Student
			tx.Find(&newStudent, newStudentId)
			newStudent.ClassID = newClass.ID
			tx.Save(newStudent)
		}
		return nil
	})
	if err != nil {
		return &error2.WebError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}
	return nil
}
