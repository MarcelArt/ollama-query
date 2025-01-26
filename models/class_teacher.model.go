package models

import "gorm.io/gorm"

const classTeacherTableName = "class_teachers"

type ClassTeacher struct {
	gorm.Model
	// Insert your fields here
	ClassID   uint `json:"classId" gorm:"not null"`
	TeacherID uint `json:"teacherId" gorm:"not null"`

	Class   *Class   `json:"-"`
	Teacher *Teacher `json:"-"`
}

type ClassTeacherDTO struct {
	DTO
	// Insert your fields here
	ClassID   uint `json:"classId" gorm:"not null"`
	TeacherID uint `json:"teacherId" gorm:"not null"`
}

type ClassTeacherPage struct {
	// Insert your fields here
	ClassID   uint `json:"classId" gorm:"not null"`
	TeacherID uint `json:"teacherId" gorm:"not null"`
}

func (ClassTeacherDTO) TableName() string {
	return classTeacherTableName
}
