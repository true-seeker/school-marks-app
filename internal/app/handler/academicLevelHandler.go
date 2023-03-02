package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"school-marks-app/internal/app/service"
)

type AcademicLevelHandler struct {
	service service.AcademicLevel
}

func NewAcademicLevelHandler(service service.AcademicLevel) *AcademicLevelHandler {
	return &AcademicLevelHandler{service: service}
}

func (a *AcademicLevelHandler) Get(c *gin.Context) {
	academicLevels, webErr := a.service.Get()
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, academicLevels)
}
