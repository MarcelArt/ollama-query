package api_routes

import (
	"github.com/MarcelArt/ollama-query/database"
	api_handlers "github.com/MarcelArt/ollama-query/handlers/api"
	"github.com/MarcelArt/ollama-query/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupAIRoutes(api fiber.Router) {
	h := api_handlers.NewAIHandler(repositories.NewTableRepo(database.GetDB()))

	g := api.Group("/ai")
	g.Post("/ask", h.Ask)
}
