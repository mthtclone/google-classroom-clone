package models

import (
	"googleCLS/config"
	"time"
)

type Assignment struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	CourseID uint `json:"course_id"`
	CourseSlug string `gorm:"-" json:"courseSlug"`
	Description string `json:"description"`
	UploadDate time.Time `json:"upload_date"`
	DueDate time.Time `json:"due_date"`
	Completed bool `json:"completed"`
	AssignedBy uint `json:"assigned_by"`
}

func GetAssignments() ([]Assignment, error) {
	var assignments []Assignment
	result := config.DB.Find(&assignments)
	return assignments, result.Error
}