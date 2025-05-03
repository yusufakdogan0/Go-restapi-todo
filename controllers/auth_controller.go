package controllers

import (
	"net/http"
	"todo-api/models"
	"todo-api/repositories"
	"todo-api/utils"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var credentials models.User
	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	for _, user := range repositories.Users {
		if user.Username == credentials.Username && user.Password == credentials.Password {
			token, err := utils.GenerateJWT(user.Username, user.Role)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"token": token,
			})
			return
		}
	}

	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
}
