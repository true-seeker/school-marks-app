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

	api := router.Group("api")
	{
		teacherGroup := api.Group("teacher")
		{
			var teacher = new(controllers.TeacherController)
			teacherGroup.GET("/:id", teacher.Get)
			teacherGroup.POST("", teacher.Create)
			teacherGroup.PATCH("/:id", teacher.Update)
			teacherGroup.DELETE("/:id", teacher.Delete)
		}
		academicLeverGroup := api.Group("academicLevel")
		{
			var academicLevel controllers.AcademicLeverController
			academicLeverGroup.GET("", academicLevel.GetAll)
		}
	}
	return router
}
