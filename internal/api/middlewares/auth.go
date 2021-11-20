package middlewares

import (
	"net/http"

	"github.com/Aibier/go-aml-service/pkg/crypto"
	"github.com/gin-gonic/gin"
)

// AuthRequired ...
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("authorization")
		if !crypto.ValidateToken(authorizationHeader) {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
			return
		}
		c.Next()
	}
}
