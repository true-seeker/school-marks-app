package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db2 "school-marks-app/api/db/models"
	"school-marks-app/api/db/models/validators"
	"strconv"
)

type AcademicYearController struct{}

func (a AcademicYearController) Get(c *gin.Context) {
	var academicYearModel db2.AcademicYear
	if c.Param("id") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing field \"id\""})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Field id must be integer"})
		return
	}

	class, webErr := academicYearModel.Get(uint(id))

	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, class)
	return
}

func (a AcademicYearController) Create(c *gin.Context) {
	var academicYear db2.AcademicYear
	if validationErr := c.BindJSON(&academicYear); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	if validationErr := validators.ValidateAcademicYearCreate(academicYear); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := academicYear.Create()
	if webErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (a AcademicYearController) Update(c *gin.Context) {
	var academicYear db2.AcademicYear
	if c.Param("id") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing field \"id\""})
		return
	}
	if validationErr := c.BindJSON(&academicYear); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "application/json data is required"})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Field id must be integer"})
		return
	}
	academicYear.ID = uint(id)

	if validationErr := validators.ValidateAcademicYearUpdate(academicYear); validationErr != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": validationErr.Error()})
		return
	}

	newTeacher, webErr := academicYear.Update()
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, newTeacher)
}

func (a AcademicYearController) Delete(c *gin.Context) {
	var academicYear db2.AcademicYear
	if c.Param("id") == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Missing field \"id\""})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Field id must be integer"})
		return
	}

	webErr := academicYear.Delete(uint(id))
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
