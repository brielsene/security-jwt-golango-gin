package routes

import (
	"security-jwt/controllers"
	"security-jwt/middlewares"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()

	// Route for login
	r.POST("/login", controllers.Login)

	// Protect routes with JWT middleware
	protected := r.Group("/")
	protected.Use(middlewares.AuthMiddleware())
	protected.GET("/welcome", controllers.Welcome)

	r.Run(":8080")
}
