package routes

import (
	"blog/controller"
	"blog/middleware"
	"blog/service"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine, uc controller.UserController, js service.JWTService) {
	userRoutes := router.Group("/user")
	{
		userRoutes.POST("/register", uc.RegisterUser)
		userRoutes.POST("/login", uc.LoginUser)
		userRoutes.GET("/me", middleware.Authenticate(js), uc.GetUser)
		userRoutes.DELETE("/me", middleware.Authenticate(js), uc.DeleteUser)
		userRoutes.PATCH("/nama", middleware.Authenticate(js), uc.UpdateUserNama)
		userRoutes.PATCH("/password", middleware.Authenticate(js), uc.UpdateUserPassword)
	}
}
