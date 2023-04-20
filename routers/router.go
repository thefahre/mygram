package routers

import (
	"mygram/controllers"
	"mygram/middlewares"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	r := gin.Default()

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
	return r
}
