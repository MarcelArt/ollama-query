package services

import (
	"encoding/json"
	"fmt"

	"github.com/MarcelArt/ollama-query/models"
	"github.com/MarcelArt/ollama-query/pkg/gollama"
	"github.com/MarcelArt/ollama-query/repositories"
)

type IAIService interface {
	Ask(ask models.AskRequest) (gollama.ChatResponse, error)
}

type AIService struct {
	tRepo repositories.ITableRepo
}

func NewAIService(tRepo repositories.ITableRepo) *AIService {
	return &AIService{
		tRepo: tRepo,
	}
}

func (s *AIService) Ask(ask models.AskRequest) (gollama.ChatResponse, error) {
	tables, err := s.tRepo.GetTablesWithColumns(ask.IncludedTables)
	if err != nil {
		return gollama.ChatResponse{}, err
	}

	tableJSON, err := json.Marshal(tables)
	if err != nil {
		return gollama.ChatResponse{}, err
	}

	g := gollama.New()

	systemContext := fmt.Sprintf(`
		You are a database assistant. Generate SQL queries based on user input without executing them, your response only contains sql query only. The database contains the following tables and columns with this specification in JSON format please do note some table is a bridge table:
		%s
	`, string(tableJSON))

	messages := []gollama.ChatMessage{}
	messages = append(messages, gollama.ChatMessage{
		Role:    "system",
		Content: systemContext,
	})
	messages = append(messages, gollama.ChatMessage{
		Role:    "user",
		Content: ask.Message,
	})

	queryRes, err := g.Chat(gollama.ChatRequest{
		Model:    "llama3.2",
		Messages: messages,
	})
	if err != nil {
		return gollama.ChatResponse{}, err
	}

	result, err := s.tRepo.RunQuery(queryRes.Message.Content)
	if err != nil {
		return gollama.ChatResponse{}, err
	}

	resJSON, err := json.Marshal(result)
	if err != nil {
		return gollama.ChatResponse{}, err
	}

	messages = append(messages, gollama.ChatMessage{
		Role:    "system",
		Content: "Now you will not generate sql query anymore but now you will summarize the json data user will send in a human-friendly way based on user previous question",
	})
	messages = append(messages, gollama.ChatMessage{
		Role:    "user",
		Content: string(resJSON),
	})

	summaryRes, err := g.Chat(gollama.ChatRequest{
		Model:    "llama3.2",
		Messages: messages,
	})

	return summaryRes, err
}
