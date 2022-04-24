package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"rest-api/src/rest/auth/controllers"
)

func setAuthRoute(router *gin.Engine) {
	authGroup := router.Group("/api/v1/auth")
	authController := new(controllers.AuthController)
	authGroup.GET("/users", authController.GetAllUsers)
	authGroup.POST("/register", authController.Register)
	authGroup.POST("/login", authController.Login)
}

func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	setAuthRoute(router)
	return router
}