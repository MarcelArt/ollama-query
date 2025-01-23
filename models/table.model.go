package models

type Table struct {
	TableName  string `json:"tableName"`
	ColumnName string `json:"columnName"`
	DataType   string `json:"dataType"`
}
