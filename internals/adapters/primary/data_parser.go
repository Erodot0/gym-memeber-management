package adapters

import (
	"log"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
)

type ErrorHandler struct{}

func (h *ErrorHandler) ParseData(c *fiber.Ctx, target interface{}) error {
	err := json.Unmarshal(c.Body(), target)
	if err != nil {
		log.Println("Error parsing data: ", err)
		return err
	}
	return nil
}