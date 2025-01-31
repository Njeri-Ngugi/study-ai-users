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
	InstitutionId  string    `gorm:"not null" json:"institution_id"`
	YearOfStudy    int       `gorm:"default:1" json:"year_of_study"`
	CompletionDate time.Time `gorm:"default:null" json:"completion_date"`
	CourseId       string    `gorm:"not null" json:"course_id"`
	Gender         string    `gorm:"not null" json:"gender"`
	DateOfBirth    time.Time `gorm:"default:null" json:"date_of_birth"`
}
