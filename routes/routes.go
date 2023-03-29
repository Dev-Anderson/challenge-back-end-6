package routes

import (
	"adopet/controllers"

	"github.com/gin-gonic/gin"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	main := router.Group("api/v1")
	{
		home := main.Group("home")
		{
			home.GET("/", controllers.GetHome)
		}
		tutor := main.Group("tutor")
		{
			tutor.GET("/", controllers.GetAllTutores)
			tutor.POST("/", controllers.CreateTutor)
			tutor.PATCH("/:id", controllers.AlterTutor)
			tutor.GET("/:id", controllers.GetIDTutor)
			tutor.DELETE("/:id", controllers.DeleteTutor)
		}
	}

	return router
}
