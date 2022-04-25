package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rest-api/src/models"
	"rest-api/src/rest/auth/services"
)

type AuthController struct{}

func (auth *AuthController) Register(c *gin.Context) {

	var newUser *models.User

	if err := c.BindJSON(&newUser); err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"message": "Invalid request body"})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.MinCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newUser.Password = string(hash)

	err = services.AuthService{}.Create(newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": newUser})

}

func (auth *AuthController) Login(c *gin.Context) {

	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"error": "Please input all fields"})
		return
	}

	usersService := services.AuthService{}

	user, mongoErr := usersService.FindOne(body.Email)
	if mongoErr != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Email or password is invalid."})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Email or password is invalid."})
		fmt.Println(user.Password, body.Password)
		return
	}

	token, jwtErr := user.GenerateJwtToken()
	if jwtErr != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": jwtErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
