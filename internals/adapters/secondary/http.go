package adapters

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/gofiber/fiber/v2"
)

type HttpServices struct{}

type Response struct {
	Success bool             `json:"success"`
	Message string           `json:"message"`
	Data    interface{}      `json:"data,omitempty"`
	Tokens  *entities.Tokens `json:"tokens,omitempty"`
}

func NewHttpServices() *HttpServices {
	return &HttpServices{}
}

// 200 OK
func (h *HttpServices) Success(c *fiber.Ctx, data interface{}, message string, tokens *entities.Tokens) error {
	return c.Status(fiber.StatusOK).JSON(Response{
		Success: true,
		Data:    data,
		Message: message,
		Tokens:  tokens,
	})
}

// 400 Bad Request
func (h *HttpServices) BadRequest(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusBadRequest).JSON(Response{
		Success: false,
		Message: message,
	})
}

// 401 Unauthorized
func (h *HttpServices) Unauthorized(c *fiber.Ctx, text string) error {
	message := "Unauthorized, please login first"

	// Set custom message
	if text != "" {
		message = text
	}

	return c.Status(fiber.StatusUnauthorized).JSON(Response{
		Success: false,
		Message: message,
	})
}

// 403 Forbidden
func (h *HttpServices) Forbidden(c *fiber.Ctx) error {
	return c.Status(fiber.StatusForbidden).JSON(Response{
		Success: false,
		Message: "Forbidden, you don't have permission to access this resource",
	})
}

// 404 Not Found
func (h *HttpServices) NotFound(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusNotFound).JSON(Response{
		Success: false,
		Message: message,
	})
}

// 500 Internal Server Error
func (h *HttpServices) InternalServerError(c *fiber.Ctx, message string) error {
	return c.Status(fiber.StatusInternalServerError).JSON(Response{
		Success: false,
		Message: message,
	})
}

// Response with html
func (h *HttpServices) WithFile(c *fiber.Ctx, pathToFile string) error {
	return c.SendFile(pathToFile)
}
