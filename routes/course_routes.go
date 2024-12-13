package routes

import (
	"tlic-api/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterCourseRoutes(router *gin.Engine, db *gorm.DB) {
	courseController := controllers.CourseController{DB: db}

	courses := router.Group("/courses")
	{
		courses.GET("", courseController.GetCourses)
		courses.GET("/:id", courseController.GetCourseByID)
		courses.POST("", courseController.CreateCourse)
		courses.PUT("/:id", courseController.UpdateCourse)
		courses.DELETE("/:id", courseController.DeleteCourse)
	}
}
