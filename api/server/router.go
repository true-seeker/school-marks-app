package server

import (
	"github.com/gin-gonic/gin"
	"school-marks-app/api/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/health", health.Status)

	v1 := router.Group("api")
	{
		teacherGroup := v1.Group("teacher")
		{
			teacher := new(controllers.TeacherController)
			teacherGroup.GET("/:id", teacher.Get)
			teacherGroup.POST("", teacher.Create)
			teacherGroup.PATCH("/:id", teacher.Update)
			teacherGroup.DELETE("/:id", teacher.Delete)
		}
	}
	return router
}
