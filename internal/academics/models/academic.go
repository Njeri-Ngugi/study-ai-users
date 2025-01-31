package models

import "github.com/Njeri-Ngugi/toolbox/models"

type Institution struct {
	models.Model    `gorm:"embedded"`
	InstitutionName string `gorm:"not null;index" json:"institution_name"`
	CountryOfStudy  string `gorm:"not null;default:'KE'" json:"country_of_study"`
}

type Course struct {
	models.Model          `gorm:"embedded"`
	CourseName            string `gorm:"not null;index" json:"course_name"`
	CourseCode            string `gorm:"not null" json:"course_code"`
	CourseDurationInYears int    `gorm:"default:4;not null" json:"course_duration"`
	InstitutionId         string `gorm:"not null;index" json:"institution_id"`
}
