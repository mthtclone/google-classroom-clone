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
}

func GetUsers() ([]User, error) {
	var users []User
	result := config.DB.Find(&users)
	return users, result.Error
}