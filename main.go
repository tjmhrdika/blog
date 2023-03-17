package main

import (
	"blog/config"
	"blog/controller"
	"blog/middleware"
	"blog/repository"
	"blog/routes"
	"blog/service"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	db := config.SetupDatabaseConnection()
	js := service.NewJWTService()
	ur := repository.NewUserRepository(db)
	us := service.NewUserService(ur)
	uc := controller.NewUserController(us, js)
	br := repository.NewBlogRepository(db)
	bs := service.NewBlogService(br)
	bc := controller.NewBlogController(bs)
	cr := repository.NewCommentRepository(db)
	cs := service.NewCommentService(cr)
	cc := controller.NewCommentController(cs)
	lr := repository.NewLikeRepository(db)
	ls := service.NewLikeService(lr)
	lc := controller.NewLikeController(ls)

	defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	routes.UserRoutes(server, uc, js)
	routes.BlogRoutes(server, bc, js)
	routes.CommentRoutes(server, cc, js)
	routes.LikeRoutes(server, lc, js)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
}
