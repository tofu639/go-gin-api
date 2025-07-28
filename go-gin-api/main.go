package main

import (
    "go-gin-api/routes"

    "github.com/gin-gonic/gin"

    // Swagger docs (will be generated)
    _ "go-gin-api/docs"
)

// @title Go Gin API
// @version 1.0
// @description This is a sample server for Go Gin API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

func main() {
    // Set Gin to release mode in production
    gin.SetMode(gin.DebugMode)

    // Create Gin router
    router := gin.Default()

    // Setup routes
    routes.SetupRoutes(router)

    // Start server
    router.Run(":8080")
}