// cmd/api/main.go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/rushabhT3/go_gin_try/internal/handlers"
    "github.com/rushabhT3/go_gin_try/internal/middlewares"
)

func main() {
    r := gin.Default()

    r.GET("/ping", handlers.Ping)

    // Apply middleware
    r.Use(middlewares.HeaderLogger())
    r.Use(middlewares.Logger())

    userRoutes := r.Group("/users")
    {
        userRoutes.GET("/", handlers.GetUsers)
        userRoutes.GET("/:id", handlers.GetUser)
        userRoutes.POST("/", handlers.CreateUser)
        userRoutes.PUT("/:id", handlers.UpdateUser)
        userRoutes.DELETE("/:id", handlers.DeleteUser)
    }

    r.Run(":8080") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
