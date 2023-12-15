package handlers

import (
	"Cinema/internal/domain/cinema/model"
	"Cinema/internal/domain/cinema/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type CinemaHandler struct {
	cinemaService service.CinemaServiceInterface
}

func NewCinemaHandler(cinemaService service.CinemaServiceInterface) *CinemaHandler {
	return &CinemaHandler{
		cinemaService: cinemaService,
	}
}

// CreateCinema
// @Summary Create a new cinema
// @Description Create a new cinema
// @Tags cinemas
// @Accept json
// @Produce json
// @Param request body model.Cinema true "Cinema object to be created"
// @Success 201 {string} string "Cinema created successfully"
// @Failure 400 {string} string "Bad Request"
// @Failure 500 {string} string "Internal Server Error"
// @Router /api/cinemas [post]
func (h *CinemaHandler) CreateCinema(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var createCinema model.CreateCinema
	if err := json.NewDecoder(r.Body).Decode(&createCinema); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.cinemaService.CreateCinema(r.Context(), createCinema); err != nil {
		http.Error(w, "Failed to create cinema", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// GetAllCinemas godoc
// @Summary Get all cinemas
// @Description Get a list of all cinemas
// @Tags cinemas
// @Produce json
// @Success 200 {array} model.Cinema
// @Failure 500 {string} Internal Server Error
// @Router /api/cinemas [get]
func (h *CinemaHandler) GetAllCinemas(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cinemas, err := h.cinemaService.GetAllCinemas(r.Context())
	if err != nil {
		http.Error(w, "Failed to get cinemas", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(cinemas)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)

}

func (h *CinemaHandler) GetCinemaByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cinemaID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid cinema ID", http.StatusBadRequest)
		return
	}

	cinema, err := h.cinemaService.GetCinemaByID(r.Context(), cinemaID)
	if err != nil {
		http.Error(w, "Failed to get cinema", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(cinema)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *CinemaHandler) UpdateCinema(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var updateCinema model.Cinema
	if err := json.NewDecoder(r.Body).Decode(&updateCinema); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.cinemaService.UpdateCinema(r.Context(), updateCinema); err != nil {
		http.Error(w, "Failed to update cinema", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *CinemaHandler) DeleteCinema(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	cinemaID, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid cinema ID", http.StatusBadRequest)
		return
	}

	if err := h.cinemaService.DeleteCinema(r.Context(), cinemaID); err != nil {
		http.Error(w, "Failed to delete cinema", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
