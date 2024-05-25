package handlers

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/gofiber/fiber/v2"
)

type MembersHandlers struct {
	Services ports.MemberServices
	Parser   ports.ParserAdapters
	Http     ports.HttpAdapters
}

func (h *MembersHandlers) CreateMember(c *fiber.Ctx) error {
	member := new(entities.Member)
	if err := h.Parser.ParseData(c, member); err != nil {
		return h.Http.BadRequest(c, "Errore nella gestione dei dati")
	}

	//Validate member
	if err := member.Validate(); err != nil {
		return h.Http.BadRequest(c, err.Error())
	}

	// Add ending date
	member.Subscription[0].AddEndDate()

	// Create member
	if err := h.Services.CreateMember(member); err != nil {
		return h.Http.InternalServerError(c, "Errore nel creare il membro")
	}

	return h.Http.Success(c, []interface{}{member}, "Membro aggiunto!")
}
