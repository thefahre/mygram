package routers

import (
	"mygram/controllers"
	"mygram/docs"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           mygram API - Fahreza
// @version         0.0.1
// @description     API endpoints for mygram

// @contact.email  thefahre@gmail.com

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" (CASE SENSITIVE!) followed by a space and JWT token. Token is obtained from login.
func StartApp() *gin.Engine {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"

	userRouter := r.Group("/users")
	{
		userRouter.POST("/register", controllers.UserRegister)
		userRouter.POST("/login", controllers.UserLogin)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middlewares.Authentication())
		photoRouter.GET("/", controllers.GetAllPhoto)
		photoRouter.GET("/:photoId", controllers.GetOnePhoto)
		photoRouter.POST("/", controllers.CreatePhoto)
		photoRouter.PUT("/:photoId", middlewares.PhotoAuthorization(), controllers.UpdatePhoto)
		photoRouter.DELETE("/:photoId", middlewares.PhotoAuthorization(), controllers.DeletePhoto)
	}

	commentRouter := r.Group("/photos/:photoId/comments")
	{
		commentRouter.Use(middlewares.Authentication())
		commentRouter.GET("/", controllers.GetAllComment)
		commentRouter.GET("/:commentId", controllers.GetOneComment)
		commentRouter.POST("/", controllers.CreateComment)
		commentRouter.PUT("/:commentId", middlewares.CommentAuthorization(), controllers.UpdateComment)
		commentRouter.DELETE("/:commentId", middlewares.CommentAuthorization(), controllers.DeleteComment)
	}

	socialmediaRouter := r.Group("/socialmedia")
	{
		socialmediaRouter.Use(middlewares.Authentication())
		socialmediaRouter.GET("/", controllers.GetAllSocialmedia)
		socialmediaRouter.GET("/:socialmediaId", controllers.GetOneSocialmedia)
		socialmediaRouter.POST("/", controllers.CreateSocialmedia)
		socialmediaRouter.PUT("/:socialmediaId", middlewares.SocialmediaAuthorization(), controllers.UpdateSocialmedia)
		socialmediaRouter.DELETE("/:socialmediaId", middlewares.SocialmediaAuthorization(), controllers.DeleteSocialmedia)
	}
	// r.GET("swagger/+any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	return r
}
