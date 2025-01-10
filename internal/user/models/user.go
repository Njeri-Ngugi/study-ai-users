package models

import (
	"github.com/Njeri-Ngugi/toolbox/models"
	"time"
)

type User struct {
	models.Model   `gorm:"embedded"`
	Firstname      string    `gorm:"not null" json:"firstname"`
	Lastname       string    `gorm:"not null" json:"lastname"`
	Username       string    `gorm:"unique,not null" json:"username"`
	Password       []byte    `gorm:"not null" json:"-"`
	Email          string    `gorm:"unique,not null" json:"email"`
	Institution    string    `gorm:"not null" json:"institution"`
	YearOfStudy    int       `gorm:"default:1" json:"year_of_study"`
	CompletionDate time.Time `gorm:"default:null" json:"completion_date"`
	CourseName     string    `gorm:"not null" json:"course_name"`
	Gender         string    `gorm:"not null" json:"gender"`
	DateOfBirth    time.Time `gorm:"default:null" json:"date_of_birth"`
}
