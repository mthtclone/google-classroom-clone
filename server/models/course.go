package models

import "googleCLS/config"

type Course struct {
	ID uint `gorm:"primaryKey" json:"id"`
	Name string `json:"name"`
	Slug string `json:"slug"`
	Section string `json:"section"`
	TeacherID uint  `json:"teacherId"`
	Teacher   User  `gorm:"foreignKey:TeacherID" json:"teacher"`

	Users []User `gorm:"many2many:course_users;" json:"users"`
}

func GetCourses() ([]Course, error) {
	var courses []Course
	result := config.DB.Preload("Users").Preload("Teacher").Find(&courses)
	return courses, result.Error
}