package middleware

import (
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// JWT secret key (should be the same as in userController.go)
var jwtKey = []byte("secret_key") // Replace this with a secure key

// Claims structure for JWT token
type Claims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

// JWT middleware to protect routes
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the Authorization header
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			c.Abort()
			return
		}

		// Token should be in the format: "Bearer <token>"
		tokenString := strings.Split(authHeader, "Bearer ")[1]

		// Parse and validate the JWT token
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil // Use the jwtKey declared above
		})

		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		// Proceed with the request
		claims, _ := token.Claims.(*Claims)
		c.Set("user_id", claims.UserID)
		c.Next()
	}
}
