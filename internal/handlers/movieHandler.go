package handlers

import (
	"Cinema/internal/domain/movie/model"
	"Cinema/internal/domain/movie/service"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type MovieHandler struct {
	movieService service.MovieServiceInterface
}

func NewMovieHandler(movieService service.MovieServiceInterface) *MovieHandler {
	return &MovieHandler{
		movieService: movieService,
	}
}

func (h *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var movie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.movieService.CreateMovie(r.Context(), movie); err != nil {
		http.Error(w, "Failed to create movie", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *MovieHandler) GetAllMovies(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	movies, err := h.movieService.GetAllMovies(r.Context())
	if err != nil {
		http.Error(w, "Failed to get movies", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *MovieHandler) GetMovieByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	movie, err := h.movieService.GetMovieByID(r.Context(), id)
	if err != nil {
		http.Error(w, "Failed to get movie", http.StatusInternalServerError)
		return
	}

	response, err := json.Marshal(movie)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
}

func (h *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	var updatedMovie model.Movie
	if err := json.NewDecoder(r.Body).Decode(&updatedMovie); err != nil {
		http.Error(w, "Failed to decode request body", http.StatusBadRequest)
		return
	}

	if err := h.movieService.UpdateMovie(r.Context(), updatedMovie); err != nil {
		http.Error(w, "Failed to update movie", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id, err := strconv.Atoi(ps.ByName("id"))
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	if err := h.movieService.DeleteMovie(r.Context(), id); err != nil {
		http.Error(w, "Failed to delete movie", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
