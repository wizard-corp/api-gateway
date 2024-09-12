package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/wizard-corp/api-gateway/src/jwttoken"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := jwttoken.IsAuthorized(authToken, secret)
			if authorized {
				userID, err := jwttoken.ExtractIDFromToken(authToken, secret)
				log.Println("2")
				if err != nil {
					c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
					c.Abort()
					return
				}
				c.Set("x-user-id", userID)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Not authorized"})
		c.Abort()
	}
}
