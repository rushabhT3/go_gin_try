package routes

import (
	"gin-auth/controllers"
	"gin-auth/middlewares"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/signup", controllers.SignUp)
	r.POST("/signin", controllers.SignIn)

	// ^ If you create a route group with r.Group("/"), it will not add an extra / to the path. The / is simply the root path.
	// ^ The resulting route will be /profile, not //profile.
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	auth.GET("/profile", controllers.GetProfile)

	return r
}
