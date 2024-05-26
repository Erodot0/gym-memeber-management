package handlers

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/utils"
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

func (h *MembersHandlers) GetMembers(c *fiber.Ctx) error {
	members, err := h.Services.GetAllMembers()
	if err != nil {
		return h.Http.InternalServerError(c, "Errore nel recuperare i membri")
	}

	return h.Http.Success(c, members, "Membri recuperati")
}

func (h *MembersHandlers) DeleteMember(c *fiber.Ctx) error {
	member := new(entities.Member)
	memberId := c.Params("id")

	// conver string to int
	id, err := utils.StringToUint(memberId)
	if err != nil {
		return h.Http.BadRequest(c, "id non valido")
	}

	member.ID = id
	if err := h.Services.GetMemberById(member); err != nil {
		return h.Http.NotFound(c, "Membero non trovato")
	}

	if err := h.Services.DeleteMember(member); err != nil {
		return h.Http.InternalServerError(c, "Errore nel eliminare il membro")
	}

	return h.Http.Success(c, nil, "Membro eliminato")
}
