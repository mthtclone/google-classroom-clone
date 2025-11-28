package models

import "googleCLS/config"

type Course struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Section string `json:"section"`
}

func GetCourses() ([]Course, error) {
	var courses []Course
	result := config.DB.Find(&courses)
	return courses, result.Error
}