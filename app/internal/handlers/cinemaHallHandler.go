package handlers

import (
	"Cinema/internal/domain/cinemaHall/model"
	"Cinema/internal/domain/cinemaHall/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CinemaHallHandler struct {
	cinemaHallService service.CinemaHallServiceInterface
}

func NewCinemaHallHandler(cinemaHallService service.CinemaHallServiceInterface) *CinemaHallHandler {
	return &CinemaHallHandler{
		cinemaHallService: cinemaHallService,
	}
}

func (h *CinemaHallHandler) CreateCinemaHall(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var cinemaHall model.CinemaHall
	if err := json.NewDecoder(r.Body).Decode(&cinemaHall); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	createCinemaHall := model.NewCreateCinemaHall(
		cinemaHall.ID,
		cinemaHall.Name,
		cinemaHall.Capacity,
		cinemaHall.Class,
		cinemaHall.CinemaId,
	)

	if err := h.cinemaHallService.CreateCinemaHall(r.Context(), createCinemaHall); err != nil {
		http.Error(w, "Failed to create cinema hall", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *CinemaHallHandler) GetCinemaHallByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid cinema hall ID", http.StatusBadRequest)
		return
	}

	cinemaHall, err := h.cinemaHallService.GetCinemaHallByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to get cinema hall", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(cinemaHall)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *CinemaHallHandler) UpdateCinemaHall(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid cinema hall ID", http.StatusBadRequest)
		return
	}

	var updatedCinemaHall model.CinemaHall
	if err := json.NewDecoder(r.Body).Decode(&updatedCinemaHall); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	updatedCinemaHall.ID = id // Ensure the ID is set to the correct value

	if err := h.cinemaHallService.UpdateCinemaHall(r.Context(), updatedCinemaHall); err != nil {
		http.Error(w, "Failed to update cinema hall", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CinemaHallHandler) GetAllCinemaHalls(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cinemaID, err := strconv.Atoi(ps.ByName("cinema_id"))
	if err != nil {
		http.Error(w, "Invalid cinema ID", http.StatusBadRequest)
		return
	}

	cinemaHalls, err := h.cinemaHallService.GetAllCinemaHalls(r.Context(), cinemaID)
	if err != nil {
		http.Error(w, "Failed to get cinema halls", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(cinemaHalls)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *CinemaHallHandler) DeleteCinemaHall(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	hallID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid cinema hall ID", http.StatusBadRequest)
		return
	}

	if err := h.cinemaHallService.DeleteCinemaHall(r.Context(), hallID); err != nil {
		http.Error(w, "Failed to delete cinema hall", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
