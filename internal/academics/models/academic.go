package models

import "github.com/Njeri-Ngugi/toolbox/models"

type Institution struct {
	models.Model    `gorm:"embedded"`
	InstitutionName string `json:"institution_name"`
	CountryCode     string `json:"country_code"`
}

type Course struct {
	models.Model   `gorm:"embedded"`
	CourseName     string `gorm:"unique,not null" json:"course_name"`
	CourseCode     string `gorm:"unique,not null" json:"course_code"`
	CourseDuration int    `gorm:"default:4;not null" json:"course_duration"`
	InstitutionId  string `gorm:"not null" json:"institution_id"`
}

type Unit struct {
	models.Model `gorm:"embedded"`
	UnitName     string `gorm:"unique,not null" json:"unit_name"`
	CourseId     string `gorm:"unique,not null" json:"course_id"`
}
