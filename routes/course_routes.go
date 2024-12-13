package routes

import (
	"tlic-api/controllers"
	"tlic-api/middlewares"
	"tlic-api/services"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCourseRoutes(router *gin.Engine, db *gorm.DB) {
	courseService := &services.CourseService{DB: db}
	courseController := controllers.CourseController{Service: courseService}

	// Add middleware for authentication and logging
	courses := router.Group("/courses")
	courses.Use(
		// middlewares.AuthMiddleware(), // apply auth to all ruotes in group
		middlewares.LoggingMiddleware(),
	)
	{
		courses.GET("", courseController.GetCourses)
		courses.GET("/:id", courseController.GetCourseByID)

		// apply auth to specific routes
		courses.POST("", middlewares.AuthMiddleware(), courseController.CreateCourse)
		courses.PUT("/:id", middlewares.AuthMiddleware(), courseController.UpdateCourse)
		courses.DELETE("/:id", middlewares.AuthMiddleware(), courseController.DeleteCourse)
	}
}
