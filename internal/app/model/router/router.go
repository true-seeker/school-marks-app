package router

import (
	"github.com/gin-gonic/gin"
	"school-marks-app/helpers"
	"school-marks-app/internal/app/handler"
	"school-marks-app/internal/app/repository"
	"school-marks-app/internal/app/service"
)

func New(r *gin.Engine) *gin.Engine {
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	api := r.Group("api")
	{
		teacherRepo := repository.NewTeacherRepository(helpers.GetConnectionOrCreateAndGet())
		teacherService := service.NewTeacherService(teacherRepo)
		teacherHandler := handler.NewTeacherController(teacherService)

		teacherGroup := api.Group("teacher")
		{
			teacherGroup.GET("/:id", teacherHandler.GetById)
			teacherGroup.POST("", teacherHandler.Create)
			teacherGroup.PATCH("/:id", teacherHandler.Update)
			teacherGroup.DELETE("/:id", teacherHandler.Delete)
		}

		academicLevelRepo := repository.NewAcademicLevelRepository(helpers.GetConnectionOrCreateAndGet())
		academicLevelService := service.NewAcademicLevelService(academicLevelRepo)
		academicLevelHandler := handler.NewAcademicLevelHandler(academicLevelService)

		academicLevelGroup := api.Group("academicLevel")
		{
			academicLevelGroup.GET("", academicLevelHandler.Get)
		}

		schoolClassRepo := repository.NewSchoolClassRepository(helpers.GetConnectionOrCreateAndGet())
		schoolClassService := service.NewSchoolClassService(schoolClassRepo)
		schoolClassHandler := handler.NewSchoolClassHandler(schoolClassService)

		schoolClassGroup := api.Group("schoolClass")
		{
			schoolClassGroup.GET("", schoolClassHandler.Get)
		}

		academicYearRepo := repository.NewAcademicYearRepository(helpers.GetConnectionOrCreateAndGet())
		academicYearService := service.NewAcademicYearService(academicYearRepo)
		academicYearHandler := handler.NewAcademicYearHandler(academicYearService)

		academicYearGroup := api.Group("academicYear")
		{
			academicYearGroup.GET("/:id", academicYearHandler.GetById)
			academicYearGroup.POST("", academicYearHandler.Create)
			academicYearGroup.PATCH("/:id", academicYearHandler.Update)
			academicYearGroup.DELETE("/:id", academicYearHandler.Delete)
		}

		studentRepo := repository.NewStudentRepository(helpers.GetConnectionOrCreateAndGet())
		classRepo := repository.NewClassRepository(helpers.GetConnectionOrCreateAndGet())
		classService := service.NewClassService(classRepo, studentRepo, teacherService, academicYearService, schoolClassService)
		classHandler := handler.NewClassHandler(classService, academicYearService, schoolClassService)

		classGroup := api.Group("class")
		{
			classGroup.GET("/:id", classHandler.GetById)
			classGroup.POST("", classHandler.Create)
			classGroup.PATCH("/:id", classHandler.Update)
			classGroup.DELETE("/:id", classHandler.Delete)
			classGroup.POST("/bulkCreate", classHandler.BulkCreate)
			classGroup.POST("/:id/transfer", classHandler.Transfer)
		}

		studentService := service.NewStudentService(studentRepo, classService)
		studentHandler := handler.NewStudentHandler(studentService)

		studentGroup := api.Group("student")
		{
			studentGroup.GET("/:id", studentHandler.GetById)
			studentGroup.POST("", studentHandler.Create)
			studentGroup.PATCH("/:id", studentHandler.Update)
			studentGroup.DELETE("/:id", studentHandler.Delete)
			studentGroup.POST("/bulkCreate", studentHandler.BulkCreate)
		}
	}
	return r
}
