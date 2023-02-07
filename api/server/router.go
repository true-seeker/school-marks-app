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
		userGroup := v1.Group("user")
		{
			user := new(controllers.UserController)
			userGroup.GET("/:id", user.Retrieve)
		}
	}
	return router

}
