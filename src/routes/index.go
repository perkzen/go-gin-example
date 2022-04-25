package routes

import (
	"github.com/gin-gonic/gin"
	_ "github.com/gin-gonic/gin"
	"rest-api/src/middlewares"
	a "rest-api/src/rest/auth/controllers"
	t "rest-api/src/rest/todos/controllers"
	u "rest-api/src/rest/users/controllers"
)

func setAuthRoute(router *gin.Engine) {
	authGroup := router.Group("/api/v1/auth")
	authController := new(a.AuthController)
	authGroup.POST("/register", authController.Register)
	authGroup.POST("/login", authController.Login)
}

func setUserRoute(router *gin.Engine) {
	usersGroup := router.Group("/api/v1/users")
	userController := new(u.UserController)
	usersGroup.Use(middlewares.Authentication())
	usersGroup.GET("/", userController.GetAllUsers)
	usersGroup.GET("/profile", userController.GetProfile)
}

func setTodoRoute(router *gin.Engine) {
	todoGroup := router.Group("/api/v1/todo")
	todoController := new(t.TodoController)
	todoGroup.POST("/", todoController.Create)
	todoGroup.GET("/:id", todoController.GetTodo)
	todoGroup.PUT("/toggle/:id", todoController.Toggle)
}

func InitRoute() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	setAuthRoute(router)
	setUserRoute(router)
	setTodoRoute(router)
	return router
}
