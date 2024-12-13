package controllers

import (
	"net/http"
	"strconv"
	"tlic-api/models"
	"tlic-api/services"
	"tlic-api/utils"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	Service *services.CourseService
}

// GET /courses
func (cc *CourseController) GetCourses(c *gin.Context) {
	courses, err := cc.Service.GetAllCourses()
	if err != nil {
		utils.RespondError(c, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondSuccess(c, http.StatusOK, courses)
}

// GET /courses/:id
func (cc *CourseController) GetCourseByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid course ID")
		return
	}
	course, err := cc.Service.GetCourseByID(id)
	if err != nil {
		utils.RespondError(c, http.StatusNotFound, "Course not found")
		return
	}
	utils.RespondSuccess(c, http.StatusOK, course)
}

// POST /courses
func (cc *CourseController) CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input")
		return
	}
	if err := cc.Service.CreateCourse(&course); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to create course")
		return
	}
	utils.RespondSuccess(c, http.StatusCreated, course)
}

// PUT /courses/:id
func (cc *CourseController) UpdateCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid course ID")
		return
	}
	var updatedCourse models.Course
	if err := c.ShouldBindJSON(&updatedCourse); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid input")
		return
	}
	course, err := cc.Service.UpdateCourse(id, &updatedCourse)
	if err != nil {
		if err.Error() == "Course not found" {
			utils.RespondError(c, http.StatusNotFound, err.Error())
			return
		}
		utils.RespondError(c, http.StatusInternalServerError, "Failed to update course")
		return
	}
	utils.RespondSuccess(c, http.StatusOK, course)
}

// DELETE /courses/:id
func (cc *CourseController) DeleteCourse(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.RespondError(c, http.StatusBadRequest, "Invalid course ID")
		return
	}
	if err := cc.Service.DeleteCourse(id); err != nil {
		utils.RespondError(c, http.StatusInternalServerError, "Failed to delete course")
		return
	}
	utils.RespondSuccess(c, http.StatusOK, gin.H{"message": "Course deleted", "course_id": id})
}
