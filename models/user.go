
package models

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
    Role     string `json:"role"` // "user" or "admin"
}
