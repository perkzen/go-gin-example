package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/src/models"
	"rest-api/src/rest/todos/services"
)

type TodoController struct{}

func (todoController TodoController) Create(c *gin.Context) {
	user := c.MustGet("user").(*models.User)

	var newTodo *models.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"message": "Invalid request body"})
		return
	}

	newTodo.User = user.Name

	err := services.TodoService{}.Create(newTodo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTodo)
}

func (todoController TodoController) Toggle(c *gin.Context) {
	id := c.Param("id")
	user := c.MustGet("user").(*models.User)

	foundTodo, err := services.TodoService{}.FindTodo(id)

	if foundTodo.User != user.Name {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "This is not your todo"})
		return
	}

	todo, err := services.TodoService{}.ToggleCompleted(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (todoController TodoController) GetTodo(c *gin.Context) {
	id := c.Param("id")
	user := c.MustGet("user").(*models.User)

	todo, err := services.TodoService{}.FindTodo(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Item not found"})
		return
	}

	if todo.User != user.Name {
		c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "This is not your todo"})
		return
	}

	c.JSON(http.StatusOK, todo)
}

func (todoController TodoController) GetUserTodos(c *gin.Context) {
	user := c.MustGet("user").(*models.User)
	todos, err := services.TodoService{}.GetTodos(user.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todos)
}
