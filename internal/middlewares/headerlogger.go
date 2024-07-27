// internal/middlewares/headerlogger.go
package middlewares

import (
    "fmt"
    "github.com/gin-gonic/gin"
)

// HeaderLogger is a middleware that logs the request headers
func HeaderLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        fmt.Println("Request Headers:")
        for key, values := range c.Request.Header {
            for _, value := range values {
                fmt.Printf("%s: %s\n", key, value)
            }
        }
        c.Next()  // Continue to the next middleware or handler
    }
}
