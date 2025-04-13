package main

import (
	"legalmatch-api/config"
	"legalmatch-api/middlewares"
	"legalmatch-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {

	config.ConnectDB()
	config.LoadEnv()
	r := gin.Default()

	middlewares.Apply(r)

	routes.RegisterAllRoutes(r)
	// config.DB.AutoMigrate(&models.User{}, &models.RefreshToken{})

	r.Run(":8080")

	// r.GET("/ping", func(ctx *gin.Context) {
	// 	ctx.JSON(200, gin.H{
	// 		"message": "pong",
	// 	})
	// })

	// r.Run(":8080")

}
