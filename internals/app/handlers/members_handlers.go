package handlers

import (
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/entities"
	"github.com/Erodot0/gym-memeber-management/internals/app/domains/ports"
	"github.com/Erodot0/gym-memeber-management/internals/app/tools/utils"
	"github.com/gofiber/fiber/v2"
)

type MembersHandlers struct {
	parser         ports.ParserAdapters
	http           ports.HttpAdapters
	memberServices ports.MemberServices
}

func NewMembersHandlers(parser ports.ParserAdapters, http ports.HttpAdapters, services ports.MemberServices) *MembersHandlers {
	return &MembersHandlers{
		parser:         parser,
		http:           http,
		memberServices: services,
	}
}

// CreateMember handles the creation of a new member.
func (h *MembersHandlers) CreateMember(c *fiber.Ctx) error {
	member := new(entities.Member)
	if err := h.parser.ParseData(c, member); err != nil {
		return h.http.BadRequest(c, "Errore nella gestione dei dati")
	}

	//Validate member
	if err := member.Validate(); err != nil {
		return h.http.BadRequest(c, err.Error())
	}

	// Add ending date
	member.Subscription[0].AddEndDate()

	// Create member
	if err := h.memberServices.CreateMember(member); err != nil {
		return h.http.InternalServerError(c, "Errore nel creare il membro")
	}

	return h.http.Success(c, []interface{}{member}, "Membro aggiunto!", nil)
}

// UpdateMember updates a member in the database.
func (h *MembersHandlers) UpdateMember(c *fiber.Ctx) error {
	updatedMember := new(entities.UpdateMember)
	if err := h.parser.ParseData(c, updatedMember); err != nil {
		return h.http.BadRequest(c, "Errore nella gestione dei dati")
	}

	// get member from fiber locals
	member := utils.GetLocalMember(c)

	// update user
	if err := h.memberServices.UpdateMember(member.ID, updatedMember); err != nil {
		return h.http.InternalServerError(c, "Errore nell'aggiornare il membro")
	}

	return h.http.Success(c, []interface{}{updatedMember}, "Membro aggiornato", nil)
}

// GetMembers retrieves all members from the database.
func (h *MembersHandlers) GetMembers(c *fiber.Ctx) error {
	members, err := h.memberServices.GetAllMembers()
	if err != nil {
		return h.http.InternalServerError(c, "Errore nel recuperare i membri")
	}

	return h.http.Success(c, members, "Membri recuperati", nil)
}

// GetMemberById retrieves a member by their ID from the database.
func (h *MembersHandlers) GetMemberById(c *fiber.Ctx) error {
	// Get member from fiber locals
	member := utils.GetLocalMember(c)

	return h.http.Success(c, []interface{}{member}, "Membro recuperato", nil)
}

// DeleteMember deletes a member from the database.
func (h *MembersHandlers) DeleteMember(c *fiber.Ctx) error {
	// Get member from fiber locals
	member := utils.GetLocalMember(c)

	// Delete member
	if err := h.memberServices.DeleteMember(member.ID); err != nil {
		return h.http.InternalServerError(c, "Errore nel eliminare il membro")
	}

	return h.http.Success(c, nil, "Membro eliminato", nil)
}

// CreateMemberSubscription handles the creation of a new member subscription.
func (h *MembersHandlers) CreateMemberSubscription(c *fiber.Ctx) error {
	// Get member from fiber locals
	member := utils.GetLocalMember(c)
	subscription := new(entities.Subscription)
	if err := h.parser.ParseData(c, subscription); err != nil {
		return h.http.BadRequest(c, "Errore nella gestione dei dati")
	}

	// Validate subscription
	if err := subscription.Validate(); err != nil {
		return h.http.BadRequest(c, err.Error())
	}

	// Add ending date
	subscription.AddEndDate()

	// Create subscription
	if err := h.memberServices.CreateMemberSubscription(member.ID, subscription); err != nil {
		return h.http.InternalServerError(c, "Errore nel creare l'iscrizione")
	}

	return h.http.Success(c, subscription, "Iscrizione creata", nil)
}

// GetMemberSubscriptions retrieves all member subscriptions from the database.
func (h *MembersHandlers) GetMemberSubscriptions(c *fiber.Ctx) error {
	// Get member from fiber locals
	member := utils.GetLocalMember(c)

	// Get subrscriptions
	subscriptions, err := h.memberServices.GetAllSubscriptions(member.ID)
	if err != nil {
		return h.http.NotFound(c, "Membros non trovato")
	}

	return h.http.Success(c, subscriptions, "Iscrizioni recuperate", nil)
}

// GetMemberSubscriptionById retrieves a member subscription by their ID from the database.
func (h *MembersHandlers) GetMemberSubscriptionById(c *fiber.Ctx) error {
	// Get member from fiber locals
	member := utils.GetLocalMember(c)
	sub_id := utils.GetUintParam(c, "sub_id")

	// Get subrscription
	subscription, err := h.memberServices.GetSubscriptionById(member.ID, sub_id)
	if err != nil {
		return h.http.NotFound(c, "Iscrizione non trovata")
	}

	return h.http.Success(c, subscription, "Iscrizione recuperata", nil)
}

// UpdateMemberSubscription updates a member subscription in the database.
func (h *MembersHandlers) UpdateMemberSubscription(c *fiber.Ctx) error {
	// Get member from fiber locals
	member := utils.GetLocalMember(c)
	sub_id := utils.GetUintParam(c, "sub_id")
	subscription := new(entities.UpdateSubscription)
	if err := h.parser.ParseData(c, subscription); err != nil {
		return h.http.BadRequest(c, "Errore nella gestione dei dati")
	}

	// Validate subscription
	if err := subscription.Validate(); err != nil {
		return h.http.BadRequest(c, err.Error())
	}

	// Add ending date
	subscription.AddEndDate()

	// Update subrscription
	updatedSub, err := h.memberServices.UpdateSubscription(member.ID, sub_id, subscription)
	if err != nil {
		return h.http.NotFound(c, "Iscrizione non trovata")
	}

	return h.http.Success(c, updatedSub, "Iscrizione aggiornata", nil)
}

// DeleteMemberSubscription deletes a member subscription from the database.
func (h *MembersHandlers) DeleteMemberSubscription(c *fiber.Ctx) error {
	// Get member from fiber locals
	member := utils.GetLocalMember(c)
	sub_id := utils.GetUintParam(c, "sub_id")

	if sub_id == 0 {
		return h.http.BadRequest(c, "Specificare l'id dell'iscrizione da eliminare")
	}

	// Get subrscription
	_, err := h.memberServices.GetSubscriptionById(member.ID, sub_id)
	if err != nil {
		return h.http.NotFound(c, "Iscrizione non trovata")
	}

	// Delete subrscription
	if err := h.memberServices.DeleteSubscription(member.ID, sub_id); err != nil {
		return h.http.NotFound(c, "Iscrizione non trovata")
	}

	return h.http.Success(c, nil, "Iscrizione eliminata", nil)
}
