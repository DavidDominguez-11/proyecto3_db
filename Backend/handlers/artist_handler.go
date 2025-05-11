// handlers/artist_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"p3db/models"
	"p3db/repositories"
	"strconv"

	"github.com/gorilla/mux"
)

// ArtistHandler expone endpoints para PerfilArtista
type ArtistHandler struct {
	repo *repositories.ArtistRepository
}

// NewArtistHandler crea un nuevo handler de artistas
func NewArtistHandler(repo *repositories.ArtistRepository) *ArtistHandler {
	return &ArtistHandler{repo: repo}
}

// GetArtists devuelve todos los perfiles de artista
func (h *ArtistHandler) GetArtists(w http.ResponseWriter, r *http.Request) {
	artists, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(artists)
}

// CreateArtist crea un nuevo perfil de artista
func (h *ArtistHandler) CreateArtist(w http.ResponseWriter, r *http.Request) {
	var a models.PerfilArtista
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := h.repo.Create(&a); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(a)
}

// GetArtist devuelve un perfil por ID
func (h *ArtistHandler) GetArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	a, err := h.repo.GetByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(a)
}

// UpdateArtist modifica un perfil existente
func (h *ArtistHandler) UpdateArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	var a models.PerfilArtista
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	a.ID = id

	if err := h.repo.Update(&a); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// DeleteArtist elimina un perfil por ID
func (h *ArtistHandler) DeleteArtist(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	if err := h.repo.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
