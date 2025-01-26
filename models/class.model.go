package models

import "gorm.io/gorm"

const classTableName = "classes"

type Class struct {
	gorm.Model
	// Insert your fields here
	Name string `json:"name" gorm:"not null"`
}

type ClassDTO struct {
	DTO
	// Insert your fields here
	Name string `json:"name" gorm:"not null"`
}

type ClassPage struct {
	// Insert your fields here
	Name string `json:"name" gorm:"not null"`
}

func (ClassDTO) TableName() string {
	return classTableName
}
