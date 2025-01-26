
package api_handlers

import (
	"github.com/MarcelArt/ollama-query/models"
	"github.com/MarcelArt/ollama-query/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ClassHandler struct {
	BaseCrudHandler[models.Class, models.ClassDTO, models.ClassPage]
	repo repositories.IClassRepo
}

func NewClassHandler(repo repositories.IClassRepo) *ClassHandler {
	return &ClassHandler{
		BaseCrudHandler: BaseCrudHandler[models.Class, models.ClassDTO, models.ClassPage]{
			repo: repo,
			validator: validator.New(validator.WithRequiredStructEnabled()),
		},
		repo: repo,
	}
}

// Create creates a new class
// @Summary Create a new class
// @Description Create a new class
// @Tags Class
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Class body models.ClassDTO true "Class data"
// @Success 201 {object} models.ClassDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /class [post]
func (h *ClassHandler) Create(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Create(c)
}

// Read retrieves a list of classes
// @Summary Get a list of classes
// @Description Get a list of classes
// @Tags Class
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.ClassPage
// @Router /class [get]
func (h *ClassHandler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing class
// @Summary Update an existing class
// @Description Update an existing class
// @Tags Class
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Class ID"
// @Param Class body models.ClassDTO true "Class data"
// @Success 200 {object} models.ClassDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /class/{id} [put]
func (h *ClassHandler) Update(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Update(c)
}

// Delete deletes an existing class
// @Summary Delete an existing class
// @Description Delete an existing class
// @Tags Class
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Class ID"
// @Success 200 {object} models.Class
// @Failure 500 {object} string
// @Router /class/{id} [delete]
func (h *ClassHandler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a class by ID
// @Summary Get a class by ID
// @Description Get a class by ID
// @Tags Class
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Class ID"
// @Success 200 {object} models.Class
// @Failure 500 {object} string
// @Router /class/{id} [get]
func (h *ClassHandler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}
