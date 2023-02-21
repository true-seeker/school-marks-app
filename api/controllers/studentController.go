package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "school-marks-app/api/db/models"
	"school-marks-app/api/db/models/validators"
)

type StudentController struct{}

func (s StudentController) GetById(c *gin.Context) {
	var studentModel db.Student
	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	teacher, webErr := studentModel.GetById(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, teacher)
	return
}

func (s StudentController) Create(c *gin.Context) {
	var student db.Student
	if validationErr := c.BindJSON(&student); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}
	if validationErr := validators.ValidateStudentCreate(student); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := student.Create()
	if webErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (s StudentController) Update(c *gin.Context) {
	var student db.Student
	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	if validationErr := c.BindJSON(&student); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	student.ID = id

	if validationErr := validators.ValidateStudentUpdate(student); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := student.Update()
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (s StudentController) Delete(c *gin.Context) {
	var student db.Student
	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	webErr := student.Delete(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
