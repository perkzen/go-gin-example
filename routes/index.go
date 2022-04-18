package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"rest-api/db/controllers"
)

func setPingRoute(router *gin.Engine) {
	authController := new(controllers.AuthController)
	router.GET("/ping", authController.Ping)
}

func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	setPingRoute(router)
	return router
}
