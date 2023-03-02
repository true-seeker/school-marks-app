package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "school-marks-app/internal/app/model/entity"
	"school-marks-app/internal/app/service"
	"school-marks-app/internal/app/validator"
)

type TeacherHandler struct {
	service service.Teacher
}

func NewTeacherController(service service.Teacher) *TeacherHandler {
	return &TeacherHandler{service: service}
}

func (t *TeacherHandler) GetById(c *gin.Context) {
	id, err := validator.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	teacher, webErr := t.service.GetById(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, teacher)
	return
}

func (t *TeacherHandler) Create(c *gin.Context) {
	teacher := &db.Teacher{}
	if validationErr := c.BindJSON(&teacher); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}
	if validationErr := validator.ValidateTeacherCreate(teacher); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := t.service.Create(teacher)
	if webErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (t *TeacherHandler) Update(c *gin.Context) {
	teacher := &db.Teacher{}
	id, err := validator.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	if validationErr := c.BindJSON(&teacher); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	teacher.ID = id

	if validationErr := validator.ValidateTeacherUpdate(teacher); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := t.service.Update(teacher)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (t *TeacherHandler) Delete(c *gin.Context) {
	id, err := validator.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	webErr := t.service.Delete(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
