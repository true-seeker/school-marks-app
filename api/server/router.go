package server

import (
	"github.com/gin-gonic/gin"
	"school-marks-app/api/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

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
		academicLevelGroup := api.Group("academicLevel")
		{
			var academicLevel controllers.AcademicLevelController
			academicLevelGroup.GET("", academicLevel.GetAll)
		}

		schoolClassGroup := api.Group("schoolClass")
		{
			var schoolClass controllers.SchoolClassController
			schoolClassGroup.GET("", schoolClass.GetAll)
		}

		academicYearGroup := api.Group("academicYear")
		{
			var academicYear controllers.AcademicYearController
			academicYearGroup.GET("/:id", academicYear.Get)
			academicYearGroup.POST("", academicYear.Create)
			academicYearGroup.PATCH("/:id", academicYear.Update)
			academicYearGroup.DELETE("/:id", academicYear.Delete)
		}
	}
	return router
}
