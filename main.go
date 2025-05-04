
package main

import (
    "todo-api/controllers"
    "todo-api/middlewares"
    "github.com/gin-gonic/gin"
)

func main() {
    r := gin.Default()
    r.POST("/login", controllers.Login)

    authorized := r.Group("/", middlewares.AuthMiddleware())
    {
	    authorized.POST("/todo-lists", controllers.CreateTodoList)
    }

    r.Run(":8080")
}
