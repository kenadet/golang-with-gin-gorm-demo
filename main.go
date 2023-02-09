package main

import (
	"github.com/gin-gonic/gin"
	"github.com/kenadet/go-gin-gorm-demo/controllers"
	"github.com/kenadet/go-gin-gorm-demo/middlewares"
	"github.com/kenadet/go-gin-gorm-demo/models"
)

func main() {

	models.ConnectDataBase()

	r := gin.Default()
	r.SetTrustedProxies([]string{"127.0.0.1"})

	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
