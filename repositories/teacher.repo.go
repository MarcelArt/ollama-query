
package repositories

import (
	"github.com/MarcelArt/ollama-query/models"
	"gorm.io/gorm"
)

const teacherPageQuery = "-- Write your query here --"

type ITeacherRepo interface {
	IBaseCrudRepo[models.Teacher, models.TeacherDTO, models.TeacherPage]
}

type TeacherRepo struct {
	BaseCrudRepo[models.Teacher, models.TeacherDTO, models.TeacherPage]
}

func NewTeacherRepo(db *gorm.DB) *TeacherRepo {
	return &TeacherRepo{
		BaseCrudRepo: BaseCrudRepo[models.Teacher, models.TeacherDTO, models.TeacherPage]{
			db:        db,
			pageQuery: teacherPageQuery,
		},
	}
}
