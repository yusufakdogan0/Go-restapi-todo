package models

import "time"

type TodoList struct {
	ID         string        `json:"id"`
	Name       string        `json:"name"`
	CreatedAt  time.Time     `json:"created_at"`
	UpdatedAt  time.Time     `json:"updated_at"`
	DeletedAt  *time.Time    `json:"deleted_at,omitempty"`
	Percentage int `json:"percentage"`
	Username   string        `json:"username"`
	Steps      []TodoStep    `json:"steps"`
}

type TodoStep struct {
	ID         string      `json:"id"`
	TodoListID string      `json:"todo_list_id"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	DeletedAt  *time.Time  `json:"deleted_at,omitempty"`
	Content    string      `json:"content"`
	IsDone     bool        `json:"is_done"`
}