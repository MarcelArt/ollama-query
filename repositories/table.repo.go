package repositories

import (
	"strings"

	"github.com/MarcelArt/ollama-query/config"
	"github.com/MarcelArt/ollama-query/models"
	"gorm.io/gorm"
)

type ITableRepo interface {
	GetTables() ([]string, error)
	GetTablesWithColumns(includedTables []string) ([]models.Table, error)
	RunQuery(query string) ([]map[string]interface{}, error)
}

type TableRepo struct {
	db *gorm.DB
}

func NewTableRepo(db *gorm.DB) *TableRepo {
	return &TableRepo{
		db: db,
	}
}

func (r *TableRepo) GetTables() ([]string, error) {
	var tables []string
	err := r.db.Table("information_schema.tables").Select("table_name").Where("table_schema = ?", "public").Pluck("table_name", &tables).Error
	return tables, err
}

func (r *TableRepo) GetTablesWithColumns(includedTables []string) ([]models.Table, error) {
	query := `
		select 
			t.table_name,
			c.column_name,
			c.data_type
		from information_schema.tables t
		join information_schema.columns c on t.table_name = c.table_name 
		where t.table_schema = ? and t.table_name in ?
		order by t.table_name 
	`
	var tables []models.Table
	if err := r.db.Raw(query, config.Env.DBSchema, includedTables).Scan(&tables).Error; err != nil {
		return nil, err
	}

	return tables, nil
}

func (r *TableRepo) RunQuery(query string) ([]map[string]interface{}, error) {
	var result []map[string]interface{}
	query = strings.ReplaceAll(query, "`", "")
	if err := r.db.Raw(query).Scan(&result).Error; err != nil {
		return nil, err
	}

	return result, nil
}
