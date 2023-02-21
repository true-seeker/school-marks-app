package validators

import (
	"errors"
	db "school-marks-app/api/db/models"
)

func ValidateClassCreate(class db.Class) error {
	if class.TeacherID == 0 {
		return errors.New("field teacher_id is missing")
	}
	if class.SchoolClassId == 0 {
		return errors.New("field school_class_id is missing")
	}
	if class.YearID == 0 {
		return errors.New("field year_id is missing")
	}
	return nil
}

func ValidateClassUpdate(class db.Class) error {
	err := ValidateClassCreate(class)
	if err != nil {
		return err
	}

	if class.ID == 0 {
		return errors.New("field id is missing")
	}
	return nil
}
