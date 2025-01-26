
package repositories

import (
	"github.com/MarcelArt/ollama-query/models"
	"gorm.io/gorm"
)

const studentPageQuery = "-- Write your query here --"

type IStudentRepo interface {
	IBaseCrudRepo[models.Student, models.StudentDTO, models.StudentPage]
}

type StudentRepo struct {
	BaseCrudRepo[models.Student, models.StudentDTO, models.StudentPage]
}

func NewStudentRepo(db *gorm.DB) *StudentRepo {
	return &StudentRepo{
		BaseCrudRepo: BaseCrudRepo[models.Student, models.StudentDTO, models.StudentPage]{
			db:        db,
			pageQuery: studentPageQuery,
		},
	}
}
