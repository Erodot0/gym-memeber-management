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

func (h *MembersHandlers) UpdateMember(c *fiber.Ctx) error {
	member := new(entities.UpdateMember)
	memberId := utils.GetApiParam(c, "id")

	// conver string to int
	id, err := utils.StringToUint(memberId)
	if err != nil {
		return h.Http.BadRequest(c, "id non valido")
	}

	// parse data
	if err := h.Parser.ParseData(c, member); err != nil {
		return h.Http.BadRequest(c, "Errore nella gestione dei dati")
	}

	member.ID = id
	if err := h.Services.UpdateMember(member); err != nil {
		return h.Http.InternalServerError(c, "Errore nel aggiornare il membro")
	}

	return h.Http.Success(c, []interface{}{member}, "Membro aggiornato")
}

func (h *MembersHandlers) GetMembers(c *fiber.Ctx) error {
	members, err := h.Services.GetAllMembers()
	if err != nil {
		return h.Http.InternalServerError(c, "Errore nel recuperare i membri")
	}

	return h.Http.Success(c, members, "Membri recuperati")
}

func (h *MembersHandlers) GetMemberById(c *fiber.Ctx) error {
	memberId := c.Params("id")

	// conver string to int
	id, err := utils.StringToUint(memberId)
	if err != nil {
		return h.Http.BadRequest(c, "id non valido")
	}

	// Get member
	member, err := h.Services.GetMemberById(id)
	if err != nil {
		return h.Http.NotFound(c, "Membro non trovato")
	}

	return h.Http.Success(c, []interface{}{member}, "Membro recuperato")
}

func (h *MembersHandlers) DeleteMember(c *fiber.Ctx) error {
	memberId := c.Params("id")

	// conver string to int
	id, err := utils.StringToUint(memberId)
	if err != nil {
		return h.Http.BadRequest(c, "id non valido")
	}

	// Get member
	_, err = h.Services.GetMemberById(id)
	if err != nil {
		return h.Http.NotFound(c, "Membro non trovato")
	}

	// Delete member
	if err := h.Services.DeleteMember(id); err != nil {
		return h.Http.InternalServerError(c, "Errore nel eliminare il membro")
	}

	return h.Http.Success(c, nil, "Membro eliminato")
}

func (h *MembersHandlers) GetMemberSubscriptions(c *fiber.Ctx) error {
	memberId := c.Params("id")

	// conver string to int
	id, err := utils.StringToUint(memberId)
	if err != nil {
		return h.Http.BadRequest(c, "id non valido")
	}

	// Get subrscriptions
	subscriptions, err := h.Services.GetAllSubscriptions(id)
	if err != nil {
		return h.Http.NotFound(c, "Membros non trovato")
	}

	return h.Http.Success(c, subscriptions, "Iscrizioni recuperate")
}