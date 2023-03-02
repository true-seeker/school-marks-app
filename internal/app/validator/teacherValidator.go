package validator

import (
	"errors"
	"school-marks-app/internal/app/model/entity"
)

func ValidateTeacherCreate(teacher *entity.Teacher) error {
	if teacher.Name == "" {
		return errors.New("field name is missing")
	}
	if teacher.Surname == "" {
		return errors.New("field surname is missing")
	}
	if teacher.Patronymic == "" {
		return errors.New("field patronymic is missing")
	}
	return nil
}

func ValidateTeacherUpdate(teacher *entity.Teacher) error {
	err := ValidateTeacherCreate(teacher)
	if err != nil {
		return err
	}
	if teacher.ID == 0 {
		return errors.New("field id is missing")
	}
	return nil
}
