package models

import "gorm.io/gorm"

const teacherTableName = "teachers"

type Teacher struct {
	gorm.Model
	// Insert your fields here
	Name string `json:"name" gorm:"not null" faker:"name"`
}

type TeacherDTO struct {
	DTO
	// Insert your fields here
	Name string `json:"name" gorm:"not null"`
}

type TeacherPage struct {
	// Insert your fields here
	Name string `json:"name" gorm:"not null"`
}

func (TeacherDTO) TableName() string {
	return teacherTableName
}
