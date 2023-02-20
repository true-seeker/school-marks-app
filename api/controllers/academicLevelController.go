package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db2 "school-marks-app/api/db/models"
)

type AcademicLevelController struct{}

func (a AcademicLevelController) Get(c *gin.Context) {
	var academicLeveModel db2.AcademicLevel

	academicLevels, webErr := academicLeveModel.Get()
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, academicLevels)
}
