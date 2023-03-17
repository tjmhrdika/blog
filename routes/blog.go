package routes

import (
	"blog/controller"
	"blog/middleware"
	"blog/service"

	"github.com/gin-gonic/gin"
)

func BlogRoutes(router *gin.Engine, bc controller.BlogController, js service.JWTService) {
	blogRoutes := router.Group("/blog")
	{
		blogRoutes.POST("", middleware.Authenticate(js), bc.CreateBlog)
		blogRoutes.GET("/:id", bc.GetBlogByID)
		blogRoutes.PATCH("/:id", middleware.Authenticate(js), bc.UpdateBlog)
		blogRoutes.DELETE("/:id", middleware.Authenticate(js), bc.DeleteBlog)
	}
}
