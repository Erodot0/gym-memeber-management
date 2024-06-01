package middlewares

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/utils"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type MemberMiddlewares struct {
	Services ports.MemberServices
	Parser   ports.ParserAdapters
	Http     ports.HttpAdapters
}


func (m *MemberMiddlewares) GetMember(c *fiber.Ctx) error {
	// Get the user ID from the API local params
	memberId := utils.GetApiParam(c, "id")


	// conver string to int
	id, err := utils.StringToUint(memberId)
	if err != nil {
		return m.Http.BadRequest(c, "id non valido")
	}

	// Retrieve the user from the database
	member, err := m.Services.GetMemberById(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return m.Http.NotFound(c, "Member not found")
		}
		return err
	}

	// Print the user
	utils.SetLocals(c, "member", &member)
	
	return c.Next()
}
