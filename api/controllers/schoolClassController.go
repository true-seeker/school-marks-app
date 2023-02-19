package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db2 "school-marks-app/api/db/models"
)

type SchoolClassController struct{}

func (s SchoolClassController) GetAll(c *gin.Context) {
	var schoolClassModel db2.SchoolClass

	schoolClasses, webErr := schoolClassModel.Get()
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, schoolClasses)
}
