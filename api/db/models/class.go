package db

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"net/http"
	"school-marks-app/api/db"
	db2 "school-marks-app/api/db/dto"
	error2 "school-marks-app/api/error"
)

// Class Класс
type Class struct {
	gorm.Model
	ParentId      uint         `json:"parent_id"`
	TeacherID     uint         `json:"teacher_id"`
	Teacher       Teacher      `json:"teacher"`
	SchoolClassId uint         `json:"school_class_id"`
	SchoolClass   SchoolClass  `json:"school_class"`
	YearID        uint         `json:"year_id"`
	Year          AcademicYear `json:"year"`
	Letter        string       `json:"letter"`
}

func ValidateClassExistingEntities(class Class) *error2.WebError {
	var teacherModel Teacher
	var yearModel AcademicYear
	var schoolClassModel SchoolClass

	if _, webErr := teacherModel.GetById(class.TeacherID); webErr != nil {
		return webErr
	}

	if _, webErr := yearModel.GetById(class.YearID); webErr != nil {
		return webErr
	}

	if _, webErr := schoolClassModel.GetById(class.SchoolClassId); webErr != nil {
		return webErr
	}
	return nil
}

func (c Class) GetById(id uint) (*Class, *error2.WebError) {
	dbConnection := db.GetDB()

	var class Class

	dbConnection.Preload(clause.Associations).Preload("SchoolClass.Level").First(&class, id)
	if class.ID == 0 {
		return nil, &error2.WebError{
			Err:  errors.New(fmt.Sprintf("class with id %d does not exist", id)),
			Code: http.StatusNotFound,
		}
	}

	return &class, nil
}

func (c Class) Create() (*Class, *error2.WebError) {
	dbConnection := db.GetDB()

	if webErr := ValidateClassExistingEntities(c); webErr != nil {
		return nil, webErr
	}

	dbConnection.Create(&c)
	dbConnection.Preload(clause.Associations).Preload("SchoolClass.Level").Find(&c)
	return &c, nil
}

func (c Class) Update() (*Class, *error2.WebError) {
	dbConnection := db.GetDB()
	oldClass, webErr := c.GetById(c.ID)
	if webErr != nil {
		return nil, webErr
	}

	if webErr = ValidateClassExistingEntities(c); webErr != nil {
		return nil, webErr
	}

	c.ParentId = oldClass.ID
	c.ID = 0
	dbConnection.Delete(&oldClass, oldClass.ID)
	dbConnection.Create(&c)
	dbConnection.Model(&Student{}).Where("class_id = ?", c.ParentId).Update("class_id", c.ID)

	dbConnection.Preload(clause.Associations).Preload("SchoolClass.Level").Find(&c)

	return &c, nil
}

func (c Class) Delete(id uint) *error2.WebError {
	dbConnection := db.GetDB()

	class, webErr := c.GetById(id)
	if webErr != nil {
		return webErr
	}

	dbConnection.Delete(&class, id)

	return nil
}

func (c Class) BulkCreate(classes []Class) ([]uint, *error2.WebError) {
	validateTeacherId := make(map[uint]bool)
	validateYearId := make(map[uint]bool)
	validateSchoolClassId := make(map[uint]bool)
	var newIds []uint
	dbConnection := db.GetDB()

	err := dbConnection.Transaction(func(tx *gorm.DB) error {
		for _, class := range classes {
			if !validateTeacherId[class.TeacherID] || !validateYearId[class.YearID] || !validateSchoolClassId[class.SchoolClassId] {
				if webErr := ValidateClassExistingEntities(class); webErr != nil {
					return webErr.Err
				}
				validateTeacherId[class.TeacherID] = true
				validateYearId[class.YearID] = true
				validateSchoolClassId[class.SchoolClassId] = true
			}
			tx.Save(&class)
			newIds = append(newIds, class.ID)
		}
		return nil
	})
	if err != nil {
		return nil, &error2.WebError{
			Err:  err,
			Code: http.StatusBadRequest,
		}
	}

	return newIds, nil
}

func (c Class) Transfer(classTransferInput db2.ClassTransferInput) *error2.WebError {
	dbConnection := db.GetDB()
	var classStudents []Student
	var classStudentIds, notTransferredStudentIds map[uint]bool
	var studentModel Student
	var academicYearModel AcademicYear
	var schoolClassModel SchoolClass

	if _, webErr := academicYearModel.GetById(classTransferInput.NewAcademicYearId); webErr != nil {
		return webErr
	}
	if _, webErr := schoolClassModel.GetById(classTransferInput.NewSchoolClassId); webErr != nil {
		return webErr
	}

	dbConnection.Where("class_id = ?", c.ID).Find(&classStudents)
	for _, student := range classStudents {
		classStudentIds[student.ID] = true
	}

	for _, id := range classTransferInput.NotTransferredStudentIds {
		if _, webErr := studentModel.GetById(id); webErr != nil {
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
		if _, webErr := studentModel.GetById(id); webErr != nil {
			return webErr
		}
		notTransferredStudentIds[id] = true
	}

	err := dbConnection.Transaction(func(tx *gorm.DB) error {
		newClass := c
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
			var newStudent Student
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
