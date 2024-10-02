package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("a big mistery")

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Get the token from Authorization Header
		authHeader := context.GetHeader("Authorization")

		if authHeader == "" {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Authorization header is empty",
			})
			context.Abort()
			return
		}

		// Split the token from the Bearer scheme
		tokenString := strings.Split(authHeader, " ")[1]

		// Parsed token

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, gin.Error{}
			}

			return jwtSecret, nil
		})

		if err != nil || !token.Valid {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			context.Abort()
			return
		}

		// Extract user_id from token clainms

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			context.Set("user_id", claims["user_id"])
		} else {
			context.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid token claims",
			})
			return
		}

		context.Next()
	}
}
