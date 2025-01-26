
package api_routes

import (
	"github.com/MarcelArt/ollama-query/database"
	api_handlers "github.com/MarcelArt/ollama-query/handlers/api"
	"github.com/MarcelArt/ollama-query/middlewares"
	"github.com/MarcelArt/ollama-query/repositories"
	"github.com/gofiber/fiber/v2"
)

func SetupClassTeacherRoutes(api fiber.Router, auth *middlewares.AuthMiddleware) {
	h := api_handlers.NewClassTeacherHandler(repositories.NewClassTeacherRepo(database.GetDB()))

	g := api.Group("/class-teacher")
	g.Get("/", auth.ProtectedAPI, h.Read)
	g.Get("/:id", auth.ProtectedAPI, h.GetByID)
	g.Post("/", auth.ProtectedAPI, h.Create)
	g.Put("/:id", auth.ProtectedAPI, h.Update)
	g.Delete("/:id", auth.ProtectedAPI, h.Delete)
}
