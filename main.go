package main

import (
	"tlic-api/configs"
	"tlic-api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// Load environment variables
	configs.LoadEnv()

	// Database connection
	db := configs.GetDBConnection()

	// Setup router
	router := gin.Default()

	// Register routes
	routes.RegisterCourseRoutes(router, db)

	// Start server
	router.Run(":8080")
}
