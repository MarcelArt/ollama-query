
package repositories

import (
	"github.com/MarcelArt/ollama-query/models"
	"gorm.io/gorm"
)

const classPageQuery = "-- Write your query here --"

type IClassRepo interface {
	IBaseCrudRepo[models.Class, models.ClassDTO, models.ClassPage]
}

type ClassRepo struct {
	BaseCrudRepo[models.Class, models.ClassDTO, models.ClassPage]
}

func NewClassRepo(db *gorm.DB) *ClassRepo {
	return &ClassRepo{
		BaseCrudRepo: BaseCrudRepo[models.Class, models.ClassDTO, models.ClassPage]{
			db:        db,
			pageQuery: classPageQuery,
		},
	}
}
