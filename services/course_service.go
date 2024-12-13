package services

import (
	"errors"
	"tlic-api/models"

	"gorm.io/gorm"
)

type CourseService struct {
	DB *gorm.DB
}

// Get all courses
func (cs *CourseService) GetAllCourses() ([]models.Course, error) {
	var courses []models.Course
	if err := cs.DB.Find(&courses).Error; err != nil {
		return nil, err
	}
	return courses, nil
}

// Get course by ID
func (cs *CourseService) GetCourseByID(id int) (*models.Course, error) {
	var course models.Course
	if err := cs.DB.First(&course, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("course not found")
		}
		return nil, err
	}
	return &course, nil
}

// Create a new course
func (cs *CourseService) CreateCourse(course *models.Course) error {
	return cs.DB.Create(course).Error
}

// Update an existing course
func (cs *CourseService) UpdateCourse(id int, updatedData *models.Course) (*models.Course, error) {
	var course models.Course
	if err := cs.DB.First(&course, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("Course not found")
		}
		return nil, err
	}
	course.Title = updatedData.Title
	course.CourseNo = updatedData.CourseNo
	course.FacultyID = updatedData.FacultyID
	course.Semester = updatedData.Semester
	if err := cs.DB.Save(&course).Error; err != nil {
		return nil, err
	}
	return &course, nil
}

// Delete a course
func (cs *CourseService) DeleteCourse(id int) error {
	return cs.DB.Delete(&models.Course{}, id).Error
}
