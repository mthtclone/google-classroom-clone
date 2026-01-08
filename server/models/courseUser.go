package models

import "googleCLS/config"

type CourseUser struct {
	ID			uint 	`gorm:"primaryKey" json:"id"`
	CourseID 	uint 	`json:"course_id"`
    UserID   	uint 	`json:"user_id"`
	Role     	string 	`json:"role"`
	User User `gorm:"foreignKey:UserID" json:"user"`
}

func GetUsersByCourse() (map[uint][]User, error) {
    var courseUsers []CourseUser
    result := config.DB.Preload("User").Find(&courseUsers)
    if result.Error != nil {
        return nil, result.Error
    }

    usersByCourse := make(map[uint][]User)
    for _, cu := range courseUsers {
        usersByCourse[cu.CourseID] = append(usersByCourse[cu.CourseID], cu.User)
    }

    return usersByCourse, nil
}