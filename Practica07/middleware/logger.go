package middleware

import (
    "github.com/gin-gonic/gin"
    "log"
    "time"
)

func LoggerMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Antes del handler
        startTime := time.Now()
        method := c.Request.Method
        path := c.Request.URL.Path
        clientIp := c.ClientIP()
        log.Printf("Request: %s %s from %s", method, path, clientIp)

        c.Next()

        endTime := time.Now()
        duration := endTime.Sub(startTime)
        statusCode := c.Writer.Status()
        log.Printf("Response: %d %s in %v", statusCode, path, duration)
    }
}
