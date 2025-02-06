package middlewares

import (
    "github.com/tokhirbekdev/AuthCore/utils"
    "net/http"
    "github.com/gin-gonic/gin"
)

func AuthRequired() gin.HandlerFunc {
    return func(c *gin.Context) {
        token := c.GetHeader("Authorization")
        if token == "" {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token required"})
            c.Abort()
            return
        }

        claims, err := utils.ValidateJWT(token)
        if err != nil {
            c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
            c.Abort()
            return
        }

        c.Set("username", claims.Username)
        c.Next()
    }
}
