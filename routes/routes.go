package routes

import (
	"adopet/controllers"
	"adopet/docs"

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
		tutor := main.Group("tutor")
		{
			tutor.GET("/", controllers.GetAllTutores)
			tutor.POST("/", controllers.CreateTutor)
			tutor.PATCH("/:id", controllers.AlterTutor)
			tutor.GET("/:id", controllers.GetIDTutor)
			tutor.DELETE("/:id", controllers.DeleteTutor)
		}
		abrigo := main.Group("abrigo")
		{
			abrigo.POST("/", controllers.CreateAbrigo)
			abrigo.GET("/", controllers.GetAllAbrigo)
			abrigo.PATCH("/:id", controllers.AlterAbrigo)
			abrigo.GET("/:id", controllers.GetIDAbrigo)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
