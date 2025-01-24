package api_handlers

import (
	"encoding/json"
	"fmt"

	"github.com/MarcelArt/ollama-query/models"
	"github.com/MarcelArt/ollama-query/pkg/gollama"
	"github.com/MarcelArt/ollama-query/repositories"
	"github.com/MarcelArt/ollama-query/utils"
	"github.com/gofiber/fiber/v2"
)

type AIHandler struct {
	tRepo repositories.ITableRepo
}

func NewAIHandler(tRepo repositories.ITableRepo) *AIHandler {
	return &AIHandler{
		tRepo: tRepo,
	}
}

// Ask about database
// @Summary Ask about database
// @Description Ask about database
// @Tags Ollama
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param AskRequest body models.AskRequest true "AskRequest data"
// @Success 200 {object} string
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /ai/ask [post]
func (h *AIHandler) Ask(c *fiber.Ctx) error {
	var ask models.AskRequest
	if err := c.BodyParser(&ask); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(models.NewJSONResponse(err, ""))
	}

	tables, err := h.tRepo.GetTablesWithColumns(ask.IncludedTables)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	tableJSON, err := json.Marshal(tables)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
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
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	result, err := h.tRepo.RunQuery(queryRes.Message.Content)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	resJSON, err := json.Marshal(result)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
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
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	return c.Status(fiber.StatusOK).JSON(models.NewJSONResponse(summaryRes.Message.Content, "Success"))
}
