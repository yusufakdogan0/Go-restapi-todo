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

func GetTodoLists(c *gin.Context) {
	username := c.GetString("username")
	role := c.GetString("role")

	var result []models.TodoList

	for _, list := range repositories.TodoLists {
		// If I am admin I should see all lists otherwise I should see my non-deleted lists
		if role == "admin" || (list.Username == username && list.DeletedAt == nil) {
			total := 0
			done := 0
			for _, step := range list.Steps {
				if step.DeletedAt == nil {
					total++
					if step.IsDone {
						done++
					}
				}
			}

			if total > 0 {
				list.Percentage = (done/total) * 100
			} else {
				list.Percentage = 0
			}

			result = append(result, list)
		}
	}

	c.JSON(http.StatusOK, result)
}

func DeleteTodoList(c *gin.Context) {
	id := c.Param("id")
	username := c.GetString("username")
	role := c.GetString("role")

	for i, list := range repositories.TodoLists {
		if list.ID == id {
			// Access control
			if role != "admin" && list.Username != username {
				c.JSON(http.StatusForbidden, gin.H{"error": "Not authorized to delete this list"})
				return
			}

			// Soft-delete by setting DeletedAt
			now := time.Now()
			repositories.TodoLists[i].DeletedAt = &now
			c.JSON(http.StatusOK, gin.H{"message": "List deleted (soft)"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "List not found"})
}


// ADD STEP
func AddTodoStep(c *gin.Context) {
	var step models.TodoStep
	if err := c.ShouldBindJSON(&step); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	for i, list := range repositories.TodoLists {
		if list.ID == step.TodoListID && list.DeletedAt == nil {
			step.ID = uuid.New().String()
			now := time.Now()
			step.CreatedAt = now
			step.UpdatedAt = now
			step.IsDone = false
			step.DeletedAt = nil

			repositories.TodoLists[i].Steps = append(repositories.TodoLists[i].Steps, step)
			c.JSON(http.StatusCreated, step)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Todo list not found"})
}

// UPDATE STEP
func UpdateTodoStep(c *gin.Context) {
	id := c.Param("id")
	var updatedData struct {
		Content string `json:"content"`
		IsDone  bool   `json:"is_done"`
	}

	if err := c.ShouldBindJSON(&updatedData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
		return
	}

	for i, list := range repositories.TodoLists {
		for j, step := range list.Steps {
			if step.ID == id && step.DeletedAt == nil {
				step.Content = updatedData.Content
				step.IsDone = updatedData.IsDone
				step.UpdatedAt = time.Now()

				repositories.TodoLists[i].Steps[j] = step
				c.JSON(http.StatusOK, step)
				return
			}
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Step not found"})
}

// DELETE STEP
func DeleteTodoStep(c *gin.Context) {
	id := c.Param("id")

	for i, list := range repositories.TodoLists {
		for j, step := range list.Steps {
			if step.ID == id && step.DeletedAt == nil {
				now := time.Now()
				repositories.TodoLists[i].Steps[j].DeletedAt = &now
				c.JSON(http.StatusOK, gin.H{"message": "Step soft-deleted"})
				return
			}
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Step not found"})
}