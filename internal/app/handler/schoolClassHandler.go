package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"school-marks-app/internal/app/service"
)

type SchoolClassHandler struct {
	service.SchoolClass
}

func NewSchoolClassHandler(schoolClass service.SchoolClass) *SchoolClassHandler {
	return &SchoolClassHandler{SchoolClass: schoolClass}
}

func (s *SchoolClassHandler) Get(c *gin.Context) {
	schoolClasses, webErr := s.SchoolClass.Get()
	if webErr != nil {
		c.AbortWithStatusJSON(webErr.Code, gin.H{"message": webErr.Err.Error()})
		return
	}
	c.JSON(http.StatusOK, schoolClasses)
}
