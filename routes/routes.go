package routes

import (
    "go-gin-api/handlers"

    "github.com/gin-gonic/gin"
    swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine) {
    // Swagger documentation endpoint
    router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

    // API routes
    api := router.Group("/api/v1")
    {
        users := api.Group("/users")
        {
            users.GET("/", handlers.GetUsers)
            users.GET("/:id", handlers.GetUser)
            users.POST("/", handlers.CreateUser)
            users.PUT("/:id", handlers.UpdateUser)
            users.DELETE("/:id", handlers.DeleteUser)
        }
    }
}