package middleware

import (
	"net/http"
	"slices"

	"github.com/gin-gonic/gin"
)

func Cors(c *gin.Context) {

	allowedOrigins := []string{
		"http://localhost:5500",
		"http://localhost:5501",
		"http://localhost:5173",
	}

	origin := c.GetHeader("Origin")

	if slices.Contains(allowedOrigins, origin) {
		c.Header("Access-Control-Allow-Origin", origin)
	}

	c.Header(
		"Access-Control-Allow-Headers",
		"Content-Type, Authorization",
	)

	c.Header(
		"Access-Control-Allow-Methods",
		"GET, POST , PUT , PATCH , DELETE , OPTIONS",
	)

	if c.Request.Method == http.MethodOptions {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()
}
