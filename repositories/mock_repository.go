package repositories

import "todo-api/models"

var Users = []models.User{
	{Username: "admin", Password: "admin123", Role: "admin"},
	{Username: "user1", Password: "user123", Role: "user"},
	{Username: "user2", Password: "user456", Role: "user"},
}

var TodoLists = []models.TodoList{}
