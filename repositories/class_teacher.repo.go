
package repositories

import (
	"github.com/MarcelArt/ollama-query/models"
	"gorm.io/gorm"
)

const classTeacherPageQuery = "-- Write your query here --"

type IClassTeacherRepo interface {
	IBaseCrudRepo[models.ClassTeacher, models.ClassTeacherDTO, models.ClassTeacherPage]
}

type ClassTeacherRepo struct {
	BaseCrudRepo[models.ClassTeacher, models.ClassTeacherDTO, models.ClassTeacherPage]
}

func NewClassTeacherRepo(db *gorm.DB) *ClassTeacherRepo {
	return &ClassTeacherRepo{
		BaseCrudRepo: BaseCrudRepo[models.ClassTeacher, models.ClassTeacherDTO, models.ClassTeacherPage]{
			db:        db,
			pageQuery: classTeacherPageQuery,
		},
	}
}
