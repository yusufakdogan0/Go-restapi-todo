
package models

import "time"

type TodoStep struct {
    ID        int
    Content   string
    IsDone    bool
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt *time.Time
}

type TodoList struct {
    ID         int
    User       string
    Name       string
    CreatedAt  time.Time
    UpdatedAt  time.Time
    DeletedAt  *time.Time
    Steps      []TodoStep
}
