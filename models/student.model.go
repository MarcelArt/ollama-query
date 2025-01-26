package models

import "gorm.io/gorm"

const studentTableName = "students"

type Student struct {
	gorm.Model
	// Insert your fields here
	Name    string `json:"name" gorm:"not null" faker:"first_name"`
	ClassID uint   `json:"classId" gorm:"not null"`

	Class *Class `json:"-"`
}

type StudentDTO struct {
	DTO
	// Insert your fields here
	Name    string `json:"name" gorm:"not null" faker:"first_name"`
	ClassID uint   `json:"classId" gorm:"not null"`
}

type StudentPage struct {
	// Insert your fields here
	Name string `json:"name" gorm:"not null"`
}

func (StudentDTO) TableName() string {
	return studentTableName
}
