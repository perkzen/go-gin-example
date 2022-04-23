package controllers

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"rest-api/src/rest/auth/services"
	"rest-api/src/rest/db/models"
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
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.MinCost)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	newUser.Password = string(hash)

	err = services.UserService{}.Create(newUser)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"user": newUser})

}

func (auth *AuthController) Login(c *gin.Context) {
	type loginDto struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	var body models.User

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"error": "Please input all fields"})
		return
	}

	usersService := services.UserService{}

	user, errf := usersService.FindOne(&body)
	if errf != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": errf.Error()})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Email or password is invalid."})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true})
}
