package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/src/db/models"
	"rest-api/src/db/services"
	_ "rest-api/src/db/services"
)

type AuthController struct{}

func (auth *AuthController) GetAllUsers(c *gin.Context) {
	users, err := services.UserService{}.GetAll()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusOK, users)
}

func (auth *AuthController) Register(c *gin.Context) {

	var newUser *models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"message": "Invalid request body"})
	}

	err := services.UserService{}.Create(newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	c.JSON(http.StatusCreated, gin.H{"user": newUser})

}
