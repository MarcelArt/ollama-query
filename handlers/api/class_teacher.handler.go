
package api_handlers

import (
	"github.com/MarcelArt/ollama-query/models"
	"github.com/MarcelArt/ollama-query/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ClassTeacherHandler struct {
	BaseCrudHandler[models.ClassTeacher, models.ClassTeacherDTO, models.ClassTeacherPage]
	repo repositories.IClassTeacherRepo
}

func NewClassTeacherHandler(repo repositories.IClassTeacherRepo) *ClassTeacherHandler {
	return &ClassTeacherHandler{
		BaseCrudHandler: BaseCrudHandler[models.ClassTeacher, models.ClassTeacherDTO, models.ClassTeacherPage]{
			repo: repo,
			validator: validator.New(validator.WithRequiredStructEnabled()),
		},
		repo: repo,
	}
}

// Create creates a new class teacher
// @Summary Create a new class teacher
// @Description Create a new class teacher
// @Tags ClassTeacher
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param ClassTeacher body models.ClassTeacherDTO true "ClassTeacher data"
// @Success 201 {object} models.ClassTeacherDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /class-teacher [post]
func (h *ClassTeacherHandler) Create(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Create(c)
}

// Read retrieves a list of class teachers
// @Summary Get a list of class teachers
// @Description Get a list of class teachers
// @Tags ClassTeacher
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.ClassTeacherPage
// @Router /class-teacher [get]
func (h *ClassTeacherHandler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing class teacher
// @Summary Update an existing class teacher
// @Description Update an existing class teacher
// @Tags ClassTeacher
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "ClassTeacher ID"
// @Param ClassTeacher body models.ClassTeacherDTO true "ClassTeacher data"
// @Success 200 {object} models.ClassTeacherDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /class-teacher/{id} [put]
func (h *ClassTeacherHandler) Update(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Update(c)
}

// Delete deletes an existing class teacher
// @Summary Delete an existing class teacher
// @Description Delete an existing class teacher
// @Tags ClassTeacher
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "ClassTeacher ID"
// @Success 200 {object} models.ClassTeacher
// @Failure 500 {object} string
// @Router /class-teacher/{id} [delete]
func (h *ClassTeacherHandler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a class teacher by ID
// @Summary Get a class teacher by ID
// @Description Get a class teacher by ID
// @Tags ClassTeacher
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "ClassTeacher ID"
// @Success 200 {object} models.ClassTeacher
// @Failure 500 {object} string
// @Router /class-teacher/{id} [get]
func (h *ClassTeacherHandler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}
