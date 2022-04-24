package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/src/rest/users/services"
)

type UserController struct{}

func (userController *UserController) GetAllUsers(c *gin.Context) {
	users, err := services.UserService{}.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, users)
}
