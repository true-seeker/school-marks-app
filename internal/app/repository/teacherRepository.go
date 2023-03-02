package repository

import (
	"gorm.io/gorm"
	"school-marks-app/internal/app/model/entity"
	error2 "school-marks-app/pkg/errorHandler"
)

type Teacher interface {
	GetById(id uint) (*entity.Teacher, *error2.WebError)
	Create(teacher *entity.Teacher) (*entity.Teacher, *error2.WebError)
	Update(newTeacher *entity.Teacher, oldTeacherId uint) (*entity.Teacher, *error2.WebError)
	Delete(id uint) *error2.WebError
}

type TeacherRepository struct {
	Db *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) Teacher {
	return &TeacherRepository{Db: db}
}

func (t *TeacherRepository) GetById(id uint) (*entity.Teacher, *error2.WebError) {
	teacher := &entity.Teacher{}
	t.Db.First(&teacher, id)
	return teacher, nil
}

func (t *TeacherRepository) Create(teacher *entity.Teacher) (*entity.Teacher, *error2.WebError) {
	t.Db.Create(&teacher)
	return teacher, nil
}

func (t *TeacherRepository) Update(newTeacher *entity.Teacher, oldTeacherId uint) (*entity.Teacher, *error2.WebError) {
	t.Db.Delete(&entity.Teacher{}, oldTeacherId)
	t.Db.Create(&newTeacher)
	t.Db.Model(&entity.Class{}).Where("teacher_id = ?", newTeacher.ParentId).Update("teacher_id", newTeacher.ID)
	return newTeacher, nil
}

func (t *TeacherRepository) Delete(id uint) *error2.WebError {
	t.Db.Delete(&entity.Teacher{}, id)
	return nil
}
