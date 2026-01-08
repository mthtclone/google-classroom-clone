package models

import (
	"googleCLS/config"
	"time"
)

type User struct {
	ID	uint `gorm:"primaryKey" json:"id"`
	NAME string `json:"name"`
	Email string `json:"email"`
	Role string `json:"role"`
	Avatar string `json:"avatar,omitempty"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Courses []Course `gorm:"many2many:course_users;" json:"courses,omitempty"`
}

func GetUsers() ([]User, error) {
	var users []User
	result := config.DB.Find(&users)
	return users, result.Error
}