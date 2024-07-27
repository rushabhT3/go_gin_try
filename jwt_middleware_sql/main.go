package main

import (
	"gin-auth/config"
	"gin-auth/models"
	"gin-auth/routes"
)

func main() {
	config.InitDB()

	// ^ This line automatically migrates the schema, creating or updating the users table based on the User model.
	config.DB.AutoMigrate(&models.User{})
	
	// ^ config.DB: The database instance.
	// ^ AutoMigrate(&models.User{}): GORM method to automatically create or update the database schema based on the User model.
	// ^ &models.User{}: A pointer to an instance of the User struct, which defines the schema.

	r := routes.SetupRouter()
	r.Run(":8080")
}

/* 
^ c *gin.Context is used within request handlers and middleware to access and manipulate request/response data.
^ gin.Engine is used to define the main application structure, routes, middleware, and start the server.
^ Use c.Get and c.Set to share data between middleware and handlers, like the userID example in the AuthMiddleware.
*/