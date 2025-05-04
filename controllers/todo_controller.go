package controllers

import (
	"net/http"
	"time"
	"todo-api/models"
	"todo-api/repositories"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func CreateTodoList(c *gin.Context) {
	var newList models.TodoList
	if err := c.ShouldBindJSON(&newList); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	username := c.GetString("username")
	newList.ID = uuid.New().String()
	newList.CreatedAt = time.Now()
	newList.UpdatedAt = time.Now()
	newList.DeletedAt = nil
	newList.Percentage = 0
	newList.Username = username 

	repositories.TodoLists = append(repositories.TodoLists, newList)

	c.JSON(http.StatusCreated, newList)
}
