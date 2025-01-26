package routes

import (
	"time"

	"github.com/MarcelArt/ollama-query/config"
	"github.com/MarcelArt/ollama-query/database"
	view_handlers "github.com/MarcelArt/ollama-query/handlers/view"
	"github.com/MarcelArt/ollama-query/middlewares"
	"github.com/MarcelArt/ollama-query/repositories"
	api_routes "github.com/MarcelArt/ollama-query/routes/api"
	view_routes "github.com/MarcelArt/ollama-query/routes/view"
	"github.com/MarcelArt/ollama-query/services"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/gofiber/swagger"
)

func SetupRoutes(app *fiber.App) {
	app.Use(cors.New())
	app.Use(logger.New())
	app.Use(limiter.New(limiter.Config{
		Max:        20,
		Expiration: 30 * time.Second,
	}))

	app.Static("/scripts", "./public/static/scripts")
	app.Static("/assets", "./public/static/assets")

	db := database.GetDB()
	tRepo := repositories.NewTableRepo(db)
	aiHandler := view_handlers.NewAIHandler(services.NewAIService(tRepo), services.NewTableService(tRepo))
	app.Get("/", aiHandler.Index)
	app.Post("/", aiHandler.Ask)

	if config.Env.ServerENV != "prod" {
		view_routes.SetupDevToolsRoutes(app)
	}

	app.Get("/swagger/*", swagger.HandlerDefault)     // default
	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
	}))

	app.Get("/metrics", monitor.New())

	authMiddleware := middlewares.NewAuthMiddleware(repositories.NewUserRepo(database.GetDB()))

	api := app.Group("/api")
	api_routes.SetupUserRoutes(api, authMiddleware)
	api_routes.SetupAuthorizedDeviceRoutes(api, authMiddleware)
	api_routes.SetupAIRoutes(api)
}
