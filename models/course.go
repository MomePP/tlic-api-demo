package models

import (
	"gorm.io/gorm"
	"time"
)

type Course struct {
	ID        uint   `gorm:"primaryKey"`
	CourseNo  string `gorm:"type:varchar(255);not null"`
	Title     string `gorm:"type:varchar(255);not null"`
	FacultyID uint   `gorm:"not null"`
	Semester  int    `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"` // Soft delete
}
