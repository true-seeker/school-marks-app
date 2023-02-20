package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db2 "school-marks-app/api/db/models"
	"school-marks-app/api/db/models/validators"
)

type ClassController struct{}

func (cl ClassController) GetById(c *gin.Context) {
	var classControllerModel db2.Class
	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	class, webErr := classControllerModel.GetById(id)

	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, class)
	return
}

func (cl ClassController) Create(c *gin.Context) {
	var class db2.Class

	if validationErr := c.BindJSON(&class); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}
	if validationErr := validators.ValidateClassCreate(class); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := class.Create()
	if webErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (cl ClassController) Update(c *gin.Context) {
	var class db2.Class

	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	if validationErr := c.BindJSON(&class); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	class.ID = id

	if validationErr := validators.ValidateClassUpdate(class); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := class.Update()
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (cl ClassController) Delete(c *gin.Context) {
	var class db2.Class
	id, err := validators.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	webErr := class.Delete(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
