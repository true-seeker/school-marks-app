package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"school-marks-app/internal/app/model/entity"
	"school-marks-app/internal/app/service"
	"school-marks-app/internal/app/validator"
)

type AcademicYearHandler struct {
	service service.AcademicYear
}

func NewAcademicYearHandler(service service.AcademicYear) *AcademicYearHandler {
	return &AcademicYearHandler{service: service}
}

func (a *AcademicYearHandler) GetById(c *gin.Context) {
	id, err := validator.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	class, webErr := a.service.GetById(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, class)
	return
}

func (a *AcademicYearHandler) Create(c *gin.Context) {
	newAcademicYear := &entity.AcademicYear{}

	if validationErr := c.BindJSON(&newAcademicYear); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	if validationErr := validator.ValidateAcademicYearCreate(newAcademicYear); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}
	academicYear, webErr := a.service.Create(newAcademicYear)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, academicYear)
}

func (a *AcademicYearHandler) Update(c *gin.Context) {
	newAcademicYear := &entity.AcademicYear{}
	id, err := validator.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		return
	}

	if validationErr := c.BindJSON(newAcademicYear); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	newAcademicYear.ID = id

	if validationErr := validator.ValidateAcademicYearUpdate(newAcademicYear); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	academicYear, webErr := a.service.Update(newAcademicYear)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, academicYear)
}

func (a *AcademicYearHandler) Delete(c *gin.Context) {
	id, err := validator.ValidateAndReturnId(c, c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	webErr := a.service.Delete(id)
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
