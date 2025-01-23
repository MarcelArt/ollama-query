package models

type AskRequest struct {
	Message        string   `json:"message"`
	IncludedTables []string `json:"includedTables"`
}
