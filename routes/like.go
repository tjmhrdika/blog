package routes

import (
	"blog/controller"
	"blog/middleware"
	"blog/service"

	"github.com/gin-gonic/gin"
)

func LikeRoutes(router *gin.Engine, lc controller.LikeController, js service.JWTService) {
	likeRoutes := router.Group("/like")
	{
		likeRoutes.POST("/:id", middleware.Authenticate(js), lc.CreateLike)
		likeRoutes.DELETE("/:id", middleware.Authenticate(js), lc.DeleteLike)

	}
}
