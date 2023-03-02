package validator

import (
	"errors"
	"school-marks-app/internal/app/model/entity"
)

func ValidateAcademicYearCreate(academicYear *entity.AcademicYear) error {
	if academicYear.Year == "" {
		return errors.New("field year is missing")
	}
	return nil
}

func ValidateAcademicYearUpdate(academicYear *entity.AcademicYear) error {
	err := ValidateAcademicYearCreate(academicYear)
	if err != nil {
		return err
	}
	if academicYear.ID == 0 {
		return errors.New("field id is missing")
	}
	return nil
}
