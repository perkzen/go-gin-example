package middlewares

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"rest-api/src/utils"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.Request.Header.Get("Token")
		if len(authHeader) == 0 {
			c.AbortWithStatusJSON(http.StatusPreconditionFailed, gin.H{"error": "Token is missing"})
			return
		}

		token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("error accurered")
			}
			secretKey := utils.EnvVar("SECRET_KEY", "")
			return []byte(secretKey), nil
		})

		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if token.Valid {
			c.Next()
		} else {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{
				"error": "Token is not valid",
			})
			return
		}
	}
}
