package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("eventhub_secret_key")

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "unauthorized, missing token",
				},
			)
			c.Abort()
			return
		}

		tokenString := strings.TrimSpace(
			strings.TrimPrefix(
				authHeader,
				"Bearer"),
		)

		token, err := jwt.Parse(
			tokenString,
			func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, errors.New("Invalid signing method")
				}
				return jwtSecret, nil
			},
		)

		if err != nil || !token.Valid {
			c.JSON(
				http.StatusUnauthorized,
				gin.H{
					"message": "invalid token",
				},
			)
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)

		if ok {
			c.Set("user_id", claims["user_id"])
			c.Set("email", claims["email"])
			c.Set("role_id", claims["role_id"])
		}
		c.Next()
	}
}
