package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "school-marks-app/api/db/models"
	"school-marks-app/api/db/models/validators"
	"strconv"
)

type TeacherController struct{}

func (t TeacherController) Get(c *gin.Context) {
	var teacherModel db.Teacher
	if c.Param("id") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing field \"id\""})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Field id must be integer"})
		return
	}

	teacher, webErr := teacherModel.Get(uint(id))
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"teacher": teacher})
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

	newTeacher, webErr := teacher.Create(teacher)
	if webErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (t TeacherController) Update(c *gin.Context) {
	var teacher db.Teacher
	if c.Param("id") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing field \"id\""})
		return
	}
	if validationErr := c.BindJSON(&teacher); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Field id must be integer"})
		return
	}
	teacher.ID = uint(id)

	if validationErr := validators.ValidateTeacherUpdate(teacher); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := teacher.Update(teacher)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (t TeacherController) Delete(c *gin.Context) {
	var teacher db.Teacher
	if c.Param("id") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing field \"id\""})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Field id must be integer"})
		return
	}

	webErr := teacher.Delete(uint(id))
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
