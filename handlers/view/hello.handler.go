package view_handlers

import (
	"github.com/MarcelArt/ollama-query/utils"
	"github.com/MarcelArt/ollama-query/views/hello"
	"github.com/gofiber/fiber/v2"
)

func HelloWorldView(c *fiber.Ctx) error {
	return utils.Render(c, hello.Show("marcel"))
}
