package handlers

import "github.com/gin-gonic/gin"

func Profile(c *gin.Context) {

	userID, _ := c.Get("user_id")
	email, _ := c.Get("email")

	c.JSON(200, gin.H{
		"message": "profile success",
		"user_id": userID,
		"email":   email,
	})
}
