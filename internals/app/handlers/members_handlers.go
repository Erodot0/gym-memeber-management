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
	updatedMember := new(entities.UpdateMember)
	if err := h.Parser.ParseData(c, updatedMember); err != nil {
		return h.Http.BadRequest(c, "Errore nella gestione dei dati")
	}

	// get member from fiber locals
	member := utils.GetLocalMember(c)

	// update user
	if err := h.Services.UpdateMember(member.ID, updatedMember); err != nil {
		return h.Http.InternalServerError(c, "Errore nell'aggiornare il membro")
	}

	return h.Http.Success(c, []interface{}{updatedMember}, "Membro aggiornato")
}

func (h *MembersHandlers) GetMembers(c *fiber.Ctx) error {
	members, err := h.Services.GetAllMembers()
	if err != nil {
		return h.Http.InternalServerError(c, "Errore nel recuperare i membri")
	}

	return h.Http.Success(c, members, "Membri recuperati")
}

func (h *MembersHandlers) GetMemberById(c *fiber.Ctx) error {
	// Get member from fiber locals
	member := utils.GetLocalMember(c)

	return h.Http.Success(c, []interface{}{member}, "Membro recuperato")
}

func (h *MembersHandlers) DeleteMember(c *fiber.Ctx) error {
	// Get member from fiber locals
	member := utils.GetLocalMember(c)

	// Delete member
	if err := h.Services.DeleteMember(member.ID); err != nil {
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

func (h *MembersHandlers) GetMemberSubscriptionById(c *fiber.Ctx) error {
	memberId := c.Params("id")
	subscriptionId := c.Params("sub_id")

	// conver string to int
	id, err := utils.StringToUint(memberId)
	if err != nil {
		return h.Http.BadRequest(c, "id membro non valido")
	}

	// conver string to int
	subId, err := utils.StringToUint(subscriptionId)
	if err != nil {
		return h.Http.BadRequest(c, "id iscrizione non valido")
	}

	// Get subrscription
	subscription, err := h.Services.GetSubscriptionById(id, subId)
	if err != nil {
		return h.Http.NotFound(c, "Membro non trovato")
	}

	return h.Http.Success(c, subscription, "Iscrizione recuperata")
}