package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/src/models"
	"rest-api/src/rest/todos/services"
)

type TodoController struct{}

func (todoController TodoController) Create(c *gin.Context) {
	var newTodo *models.Todo

	if err := c.BindJSON(&newTodo); err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"message": "Invalid request body"})
		return
	}

	err := services.TodoService{}.Create(newTodo)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, newTodo)
}

func (todoController TodoController) Toggle(c *gin.Context) {
	id := c.Param("id")

	if len(id) == 0 {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"message": "Id is missing"})
		return
	}

	todo, err := services.TodoService{}.ToggleCompleted(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, todo)

}
