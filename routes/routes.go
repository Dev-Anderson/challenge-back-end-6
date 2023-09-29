package routes

import (
	"adopet/controllers"
	"adopet/docs"
	"adopet/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func ConfigRoutes(router *gin.Engine) *gin.Engine {
	docs.SwaggerInfo.BasePath = "api/v1"
	main := router.Group("api/v1")
	{
		home := main.Group("home")
		{
			home.GET("/", controllers.GetHome)
		}
		user := main.Group("user")
		{
			user.GET("/", controllers.GetAllUser)
			user.POST("/", controllers.CreateUser)
		}
		login := main.Group("login")
		{
			login.POST("/", controllers.Login)
		}

		tutor := main.Group("tutor", middleware.Auth())
		{
			tutor.GET("/", controllers.GetAllTutores)
			tutor.POST("/", controllers.CreateTutor)
			tutor.PATCH("/:id", controllers.AlterTutor)
			tutor.GET("/:id", controllers.GetIDTutor)
			tutor.DELETE("/:id", controllers.DeleteTutor)
		}
		abrigo := main.Group("abrigo", middleware.Auth())
		{
			abrigo.POST("/", controllers.CreateAbrigo)
			abrigo.GET("/", controllers.GetAllAbrigo)
			abrigo.PATCH("/:id", controllers.AlterAbrigo)
			abrigo.GET("/:id", controllers.GetIDAbrigo)
			abrigo.DELETE("/:id", controllers.DeleteAbrigo)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
