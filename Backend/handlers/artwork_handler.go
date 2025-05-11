// handlers/artwork_handler.go
package handlers

import (
	"encoding/json"
	"net/http"
	"p3db/models"
	"p3db/repositories"
	"strconv"

	"github.com/gorilla/mux"
)

// ArtworkHandler expone endpoints para ObraArte
type ArtworkHandler struct {
	repo *repositories.ArtworkRepository
}

// NewArtworkHandler crea un nuevo handler de obras de arte
func NewArtworkHandler(repo *repositories.ArtworkRepository) *ArtworkHandler {
	return &ArtworkHandler{repo: repo}
}

// GetArtworks devuelve todas las obras de arte
func (h *ArtworkHandler) GetArtworks(w http.ResponseWriter, r *http.Request) {
	list, err := h.repo.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(list)
}

// CreateArtwork crea una nueva obra de arte
func (h *ArtworkHandler) CreateArtwork(w http.ResponseWriter, r *http.Request) {
	var a models.ObraArte
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

// GetArtwork devuelve una obra por ID
func (h *ArtworkHandler) GetArtwork(w http.ResponseWriter, r *http.Request) {
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

// UpdateArtwork modifica una obra existente
func (h *ArtworkHandler) UpdateArtwork(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	var a models.ObraArte
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

// DeleteArtwork elimina una obra por ID
func (h *ArtworkHandler) DeleteArtwork(w http.ResponseWriter, r *http.Request) {
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