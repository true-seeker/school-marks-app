package validator

import (
	"errors"
	"school-marks-app/internal/app/model/entity"
)

func ValidateStudentCreate(student *entity.Student) error {
	if student.Name == "" {
		return errors.New("field name is missing")
	}
	if student.Surname == "" {
		return errors.New("field surname is missing")
	}
	if student.Patronymic == "" {
		return errors.New("field patronymic is missing")
	}
	if student.ClassID == 0 {
		return errors.New("field class_id is missing")
	}
	return nil
}

func ValidateStudentUpdate(student *entity.Student) error {
	err := ValidateStudentCreate(student)
	if err != nil {
		return err
	}
	if student.ID == 0 {
		return errors.New("field id is missing")
	}
	return nil
}
