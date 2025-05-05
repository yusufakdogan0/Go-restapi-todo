
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
        authorized.GET("/todo-lists", controllers.GetTodoLists)
        authorized.DELETE("/todo-lists/:id", controllers.DeleteTodoList)
        authorized.POST("/todo-steps", controllers.AddTodoStep)
        authorized.PUT("/todo-steps/:id", controllers.UpdateTodoStep)
        authorized.DELETE("/todo-steps/:id", controllers.DeleteTodoStep)


    }

    r.Run(":8080")
}
