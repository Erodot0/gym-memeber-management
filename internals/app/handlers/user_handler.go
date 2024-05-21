package handlers

import (
	"fmt"

	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/gofiber/fiber/v2"
)

type UserHandlers struct {
	Parser ports.ParserAdapters
	Http ports.HttpAdapters
}

func (h *UserHandlers) CreateUser(c *fiber.Ctx) error {
	user := new(entities.User)
	if err := h.Parser.ParseData(c, user); err != nil {
		return h.Http.BadRequest(c, "Error parsing data")
	}

	fmt.Println(user)

	user.RemovePassword()
	return h.Http.Success(c, []interface{}{user}, "User created")
}