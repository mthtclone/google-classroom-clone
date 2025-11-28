package models

import "googleCLS/config"

type ResourceFile struct {
	Name string `json:"name"`
	URL string `json:"url"`
}

type Resource struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Title string `json:"title"`
	CourseID uint `json:"course_id"`
	Description string `json:"description"`
	UploadedBy uint `json:"uploaded_by"` // user.id
	AssignmentID *uint `json:"assignment_iud,omitempty"`
	Files []ResourceFile `gorm:"-" json:"files"` // ignore GORM, handle in app
}

func GetResources() ([]Resource, error) {
	var resources []Resource
	result := config.DB.Find(&resources)
	return resources, result.Error
}