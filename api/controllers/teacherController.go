package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "school-marks-app/api/db/models"
	"school-marks-app/api/db/models/validators"
)

type TeacherController struct{}

func (t TeacherController) Get(c *gin.Context) {
	var teacherModel db.Teacher
	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	teacher, webErr := teacherModel.Get(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, teacher)
	return
}

func (t TeacherController) Create(c *gin.Context) {
	var teacher db.Teacher
	if validationErr := c.BindJSON(&teacher); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}
	if validationErr := validators.ValidateTeacherCreate(teacher); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := teacher.Create()
	if webErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (t TeacherController) Update(c *gin.Context) {
	var teacher db.Teacher
	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	if validationErr := c.BindJSON(&teacher); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	teacher.ID = id

	if validationErr := validators.ValidateTeacherUpdate(teacher); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := teacher.Update()
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (t TeacherController) Delete(c *gin.Context) {
	var teacher db.Teacher
	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	webErr := teacher.Delete(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
