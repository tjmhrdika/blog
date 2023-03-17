package routes

import (
	"blog/controller"
	"blog/middleware"
	"blog/service"

	"github.com/gin-gonic/gin"
)

func CommentRoutes(router *gin.Engine, cc controller.CommentController, js service.JWTService) {
	commentRoutes := router.Group("/comment")
	{
		commentRoutes.POST("/:id", middleware.Authenticate(js), cc.CreateComment)
		commentRoutes.PUT("/:id", middleware.Authenticate(js), cc.UpdateComment)
		commentRoutes.DELETE("/:id", middleware.Authenticate(js), cc.DeleteComment)
	}
}
