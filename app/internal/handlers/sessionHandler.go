package handlers

import (
	"Cinema/internal/domain/session/model"
	"Cinema/internal/domain/session/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type SessionHandler struct {
	sessionService service.SessionServiceInterface
}

func NewSessionHandler(sessionService service.SessionServiceInterface) *SessionHandler {
	return &SessionHandler{
		sessionService: sessionService,
	}
}

func (h *SessionHandler) GetSessionByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	session, err := h.sessionService.GetSessionByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to get session", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(session)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *SessionHandler) GetSessionsByCinemaHallID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cinemaHallID, err := strconv.Atoi(ps.ByName("cinemaHallId"))
	if err != nil {
		http.Error(w, "Invalid cinema hall ID", http.StatusBadRequest)
		return
	}

	sessions, err := h.sessionService.GetSessionsByCinemaHallID(r.Context(), cinemaHallID)
	if err != nil {
		http.Error(w, "Failed to get sessions by cinema hall ID", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(sessions)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *SessionHandler) CreateSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var session model.Session
	if err := json.NewDecoder(r.Body).Decode(&session); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.sessionService.CreateSession(r.Context(), session); err != nil {
		http.Error(w, "Failed to create session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *SessionHandler) UpdateSession(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var updatedSession model.Session
	if err := json.NewDecoder(r.Body).Decode(&updatedSession); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	err := h.sessionService.UpdateSession(r.Context(), updatedSession)
	if err != nil {
		http.Error(w, "Failed to update session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *SessionHandler) DeleteSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid session ID", http.StatusBadRequest)
		return
	}

	err = h.sessionService.DeleteSession(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to delete session", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
