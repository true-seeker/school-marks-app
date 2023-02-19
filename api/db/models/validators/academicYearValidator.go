package validators

import (
	"errors"
	db "school-marks-app/api/db/models"
)

func ValidateAcademicYearCreate(academicYear db.AcademicYear) error {
	if academicYear.Year == "" {
		return errors.New("field year is missing")
	}
	return nil
}

func ValidateAcademicYearUpdate(academicYear db.AcademicYear) error {
	err := ValidateAcademicYearCreate(academicYear)
	if err != nil {
		return err
	}
	if academicYear.ID == 0 {
		return errors.New("field id is missing")
	}
	return nil
}
