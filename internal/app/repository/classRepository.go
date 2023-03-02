package repository

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"school-marks-app/internal/app/model/entity"
	db2 "school-marks-app/internal/app/model/input"
	error2 "school-marks-app/pkg/errorHandler"
)

type Class interface {
	GetById(id uint) (*entity.Class, *error2.WebError)
	Create(class *entity.Class) (*entity.Class, *error2.WebError)
	Update(class *entity.Class, oldClassId uint) (*entity.Class, *error2.WebError)
	Delete(id uint) *error2.WebError
	BulkCreate(classes *[]entity.Class) ([]uint, *error2.WebError)
	Transfer(classTransferInput db2.ClassTransferInput) *error2.WebError
}

type ClassRepository struct {
	Db *gorm.DB
}

func NewClassRepository(db *gorm.DB) Class {
	return &ClassRepository{Db: db}
}

func (c *ClassRepository) GetById(id uint) (*entity.Class, *error2.WebError) {
	class := &entity.Class{}
	c.Db.Preload(clause.Associations).Preload("SchoolClass.Level").First(&class, id)
	return class, nil
}

func (c *ClassRepository) Create(class *entity.Class) (*entity.Class, *error2.WebError) {
	c.Db.Create(&class)
	c.Db.Preload(clause.Associations).Preload("SchoolClass.Level").Find(&class)
	return class, nil
}

func (c *ClassRepository) Update(class *entity.Class, oldClassId uint) (*entity.Class, *error2.WebError) {
	c.Db.Delete(&entity.Class{}, oldClassId)
	c.Db.Create(&class)
	c.Db.Model(&entity.Student{}).Where("class_id = ?", class.ParentId).Update("class_id", class.ID)

	c.Db.Preload(clause.Associations).Preload("SchoolClass.Level").Find(&class)
	return class, nil
}

func (c *ClassRepository) Delete(id uint) *error2.WebError {
	c.Db.Delete(&entity.Class{}, id)
	return nil
}

func (c *ClassRepository) BulkCreate(classes *[]entity.Class) ([]uint, *error2.WebError) {
	var ids []uint
	for _, class := range *classes {
		c.Create(&class)
		ids = append(ids, class.ID)
	}
	return ids, nil
}

func (c *ClassRepository) Transfer(classTransferInput db2.ClassTransferInput) *error2.WebError {
	//TODO implement me
	panic("implement me")
}
