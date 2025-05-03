
package main

import (
    "todo-api/controllers"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.POST("/login", controllers.Login)
    r.Run(":8080")
}
