package routes

import (
	"golang-microservice/controllers"
	"golang-microservice/middleware"

	"github.com/gin-gonic/gin"
)

func UserRoutes(router *gin.Engine) {
	secured := router.Group("/users").Use(middleware.AuthMiddleware())

	secured.POST("", controllers.CreateUser)
	secured.GET("/:id", controllers.GetUser)
	secured.PUT("/:id", controllers.UpdateUser)
	secured.DELETE("/:id", controllers.DeleteUser)
}

func AuthRoutes(router *gin.Engine) {
	router.GET("/auth/login", middleware.OAuthLogin)
	router.GET("/auth/callback", middleware.OAuthCallback)
	router.POST("/login", controllers.Login)
}
