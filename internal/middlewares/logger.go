// internal/middlewares/logger.go
package middlewares

import (
    "fmt"
    "time"
    "github.com/gin-gonic/gin"
)

// Logger is a middleware that logs the request details
func Logger() gin.HandlerFunc {
    return func(c *gin.Context) {
        t := time.Now()

        // Set example variable
        c.Set("example", "12345")

        // Before request
        c.Next()

        // After request
        latency := time.Since(t)
        status := c.Writer.Status()

        fmt.Printf("Request: %s %s | Status: %d | Latency: %s\n",
            c.Request.Method, c.Request.URL.Path, status, latency)
    }
}
