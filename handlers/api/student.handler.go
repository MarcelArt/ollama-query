
package api_handlers

import (
	"github.com/MarcelArt/ollama-query/models"
	"github.com/MarcelArt/ollama-query/repositories"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type StudentHandler struct {
	BaseCrudHandler[models.Student, models.StudentDTO, models.StudentPage]
	repo repositories.IStudentRepo
}

func NewStudentHandler(repo repositories.IStudentRepo) *StudentHandler {
	return &StudentHandler{
		BaseCrudHandler: BaseCrudHandler[models.Student, models.StudentDTO, models.StudentPage]{
			repo: repo,
			validator: validator.New(validator.WithRequiredStructEnabled()),
		},
		repo: repo,
	}
}

// Create creates a new student
// @Summary Create a new student
// @Description Create a new student
// @Tags Student
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param Student body models.StudentDTO true "Student data"
// @Success 201 {object} models.StudentDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /student [post]
func (h *StudentHandler) Create(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Create(c)
}

// Read retrieves a list of students
// @Summary Get a list of students
// @Description Get a list of students
// @Tags Student
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param page query int false "Page"
// @Param size query int false "Size"
// @Param sort query string false "Sort"
// @Param filters query string false "Filter"
// @Success 200 {array} models.StudentPage
// @Router /student [get]
func (h *StudentHandler) Read(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Read(c)
}

// Update updates an existing student
// @Summary Update an existing student
// @Description Update an existing student
// @Tags Student
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Student ID"
// @Param Student body models.StudentDTO true "Student data"
// @Success 200 {object} models.StudentDTO
// @Failure 400 {object} string
// @Failure 500 {object} string
// @Router /student/{id} [put]
func (h *StudentHandler) Update(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Update(c)
}

// Delete deletes an existing student
// @Summary Delete an existing student
// @Description Delete an existing student
// @Tags Student
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Student ID"
// @Success 200 {object} models.Student
// @Failure 500 {object} string
// @Router /student/{id} [delete]
func (h *StudentHandler) Delete(c *fiber.Ctx) error {
	return h.BaseCrudHandler.Delete(c)
}

// GetByID retrieves a student by ID
// @Summary Get a student by ID
// @Description Get a student by ID
// @Tags Student
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param id path string true "Student ID"
// @Success 200 {object} models.Student
// @Failure 500 {object} string
// @Router /student/{id} [get]
func (h *StudentHandler) GetByID(c *fiber.Ctx) error {
	return h.BaseCrudHandler.GetByID(c)
}
