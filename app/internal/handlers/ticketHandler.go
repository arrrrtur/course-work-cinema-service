package handlers

import (
	"Cinema/internal/domain/ticket/model"
	"Cinema/internal/domain/ticket/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type TicketHandler struct {
	ticketService service.TicketServiceInterface
}

func NewTicketHandler(ticketService service.TicketServiceInterface) *TicketHandler {
	return &TicketHandler{
		ticketService: ticketService,
	}
}

func (h *TicketHandler) CreateTicket(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var ticket model.Ticket
	if err := json.NewDecoder(r.Body).Decode(&ticket); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.ticketService.CreateTicket(r.Context(), ticket); err != nil {
		http.Error(w, "Failed to create ticket", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *TicketHandler) GetTicketByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	ticket, err := h.ticketService.GetTicketByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to get ticket", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(ticket)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *TicketHandler) GetTicketsByUserID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userId, err := strconv.Atoi(ps.ByName("userId"))
	if err != nil {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	tickets, err := h.ticketService.GetTicketsByUserID(r.Context(), userId)
	if err != nil {
		http.Error(w, "Failed to get tickets", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(tickets)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *TicketHandler) UpdateTicket(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	var updatedTicket model.Ticket
	if err := json.NewDecoder(r.Body).Decode(&updatedTicket); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.ticketService.UpdateTicket(r.Context(), id, updatedTicket); err != nil {
		http.Error(w, "Failed to update ticket", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *TicketHandler) DeleteTicket(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid ticket ID", http.StatusBadRequest)
		return
	}

	if err := h.ticketService.DeleteTicket(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete ticket", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
