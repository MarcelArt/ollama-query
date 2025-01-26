
package api_handlers

import (
	"github.com/MarcelArt/ollama-query/models"
	"github.com/MarcelArt/ollama-query/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type TeacherHandler struct {
	BaseCrudHandler[models.Teacher, models.TeacherDTO, models.TeacherPage]
	repo repositories.ITeacherRepo
}

func NewTeacherHandler(repo repositories.ITeacherRepo) *TeacherHandler {
	return &TeacherHandler{
		BaseCrudHandler: BaseCrudHandler[models.Teacher, models.TeacherDTO, models.TeacherPage]{
			repo: repo,
			validator: validator.New(validator.WithRequiredStructEnabled()),
		},
		repo: repo,
	}
}

// Create creates a new teacher
// @Summary Create a new teacher
// @Description Create a new teacher
// @Tags Teacher
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Teacher body models.TeacherDTO true "Teacher data"
// @Success 201 {object} models.TeacherDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /teacher [post]
func (h *TeacherHandler) Create(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Create(c)
}

// Read retrieves a list of teachers
// @Summary Get a list of teachers
// @Description Get a list of teachers
// @Tags Teacher
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.TeacherPage
// @Router /teacher [get]
func (h *TeacherHandler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing teacher
// @Summary Update an existing teacher
// @Description Update an existing teacher
// @Tags Teacher
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Teacher ID"
// @Param Teacher body models.TeacherDTO true "Teacher data"
// @Success 200 {object} models.TeacherDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /teacher/{id} [put]
func (h *TeacherHandler) Update(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Update(c)
}

// Delete deletes an existing teacher
// @Summary Delete an existing teacher
// @Description Delete an existing teacher
// @Tags Teacher
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Teacher ID"
// @Success 200 {object} models.Teacher
// @Failure 500 {object} string
// @Router /teacher/{id} [delete]
func (h *TeacherHandler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a teacher by ID
// @Summary Get a teacher by ID
// @Description Get a teacher by ID
// @Tags Teacher
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Teacher ID"
// @Success 200 {object} models.Teacher
// @Failure 500 {object} string
// @Router /teacher/{id} [get]
func (h *TeacherHandler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}
