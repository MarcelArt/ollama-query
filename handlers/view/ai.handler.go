package view_handlers

import (
	"github.com/MarcelArt/ollama-query/models"
	"github.com/MarcelArt/ollama-query/services"
	"github.com/MarcelArt/ollama-query/utils"
	"github.com/MarcelArt/ollama-query/views/ai"
	"github.com/MarcelArt/ollama-query/views/components"
	"github.com/gofiber/fiber/v2"
)

type AIHandler struct {
	service  services.IAIService
	tService services.ITableService
}

func NewAIHandler(service services.IAIService, tService services.ITableService) *AIHandler {
	return &AIHandler{
		service:  service,
		tService: tService,
	}
}

func (h *AIHandler) Index(c *fiber.Ctx) error {
	tables, err := h.tService.GetTables()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	return utils.Render(c, ai.Index(tables))
}

func (h *AIHandler) Ask(c *fiber.Ctx) error {
	var ask models.AskRequest
	if err := c.BodyParser(&ask); err != nil {
		return utils.Render(c, components.ResponseSection(err.Error()))
	}

	summaryRes, err := h.service.Ask(ask)
	if err != nil {
		return utils.Render(c, components.ResponseSection(err.Error()))
	}

	return utils.Render(c, components.ResponseSection(summaryRes.Message.Content))

}
