package middlewares

import (
	"net/http"
	"strings"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Extract token from the Authorization header
		authHeader := c.GetHeader("Authorization")
		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing or malformed"})
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		secretKey := os.Getenv("AUTH_SECRET_KEY") // Load secret key from environment

		// Validate token (you'd use a JWT library here)
		isValid, claims := validateJWTToken(token, secretKey)
		if !isValid {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		// Add claims to context for use in handlers
		c.Set("user_id", claims["user_id"])
		c.Next() // Proceed to the next middleware or handler
	}
}

func validateJWTToken(token, secretKey string) (bool, map[string]interface{}) {
	// Pseudo JWT validation logic
	// Replace with an actual library like github.com/dgrijalva/jwt-go
	return token == secretKey, map[string]interface{}{"user_id": "123"}
}
