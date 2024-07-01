package auth

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// AuthAPIKeyMiddleware is a middleware function for Gin that validates the API key in the request header.
func AuthAPIKeyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		apiKey := c.GetHeader("X-API-Key")
		expectedAPIKey := os.Getenv("MY_API_KEY")

		if apiKey == "" || apiKey != expectedAPIKey {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
}
