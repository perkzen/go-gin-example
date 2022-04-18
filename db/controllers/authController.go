package controllers

import "github.com/gin-gonic/gin"

type AuthController struct{}

func (auth *AuthController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{"ping": "ok"})
}
