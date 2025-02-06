package routes

import (
    "github.com/tokhirbekdev/AuthCore/controllers"
    "github.com/tokhirbekdev/AuthCore/middlewares"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    router := gin.Default()

    authRoutes := router.Group("/auth")
    {
        authRoutes.POST("/register", controllers.Register)
        authRoutes.POST("/login", controllers.Login)
    }

    protectedRoutes := router.Group("/protected")
    protectedRoutes.Use(middlewares.AuthRequired())
    {
        // protected route example
        protectedRoutes.GET("/profile", func(c *gin.Context) {
            username, _ := c.Get("username")
            c.JSON(200, gin.H{"message": "Welcome " + username.(string)})
        })
    }

    return router
}
