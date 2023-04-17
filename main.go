package main

import (
	"github.com/Alfred-mk/go-jwt/controllers"
	"github.com/Alfred-mk/go-jwt/initializers"
	"github.com/Alfred-mk/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.POST("/signup", controllers.Signup)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)
	r.POST("/posts", middleware.RequireAuth, controllers.PostsCreate)
	r.GET("/posts", middleware.RequireAuth, controllers.PostsIndex)
	r.GET("/posts/:id", middleware.RequireAuth, controllers.PostsShow)

	r.Run()
}
