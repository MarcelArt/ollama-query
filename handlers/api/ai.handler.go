package api_handlers

import (
	"github.com/MarcelArt/ollama-query/models"
	"github.com/MarcelArt/ollama-query/services"
	"github.com/MarcelArt/ollama-query/utils"
	"github.com/gofiber/fiber/v2"
)

type AIHandler struct {
	service services.IAIService
}

func NewAIHandler(service services.IAIService) *AIHandler {
	return &AIHandler{
		service: service,
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

	summaryRes, err := h.service.Ask(ask)
	if err != nil {
		return c.Status(utils.StatusCodeByError(err)).JSON(models.NewJSONResponse(err, ""))
	}

	return c.Status(fiber.StatusOK).JSON(models.NewJSONResponse(summaryRes.Message.Content, "Success"))
}
