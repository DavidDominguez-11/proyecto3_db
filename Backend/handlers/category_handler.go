// handlers/category_handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "p3db/models"
    "p3db/repositories"
)

type CategoryHandler struct {
    repo *repositories.CategoryRepository
}

func NewCategoryHandler(repo *repositories.CategoryRepository) *CategoryHandler {
    return &CategoryHandler{repo: repo}
}

func (h *CategoryHandler) GetCategoryArtworks(w http.ResponseWriter, r *http.Request) {
    var filter models.CategoryArtworkFilter
    
    q := r.URL.Query()
    
    // Parsear parámetros requeridos
    categoriaID, err := strconv.Atoi(q.Get("categoria_id"))
    if err != nil || categoriaID == 0 {
        http.Error(w, "categoria_id es requerido y debe ser numérico", http.StatusBadRequest)
        return
    }
    filter.CategoriaID = categoriaID
    
    // Parsear parámetros opcionales
    if precioMax := q.Get("precio_max"); precioMax != "" {
        if max, err := strconv.ParseFloat(precioMax, 64); err == nil {
            filter.PrecioMax = max
        }
    }
    
    filter.Estado = q.Get("estado")
    
    if artistaID := q.Get("artista_id"); artistaID != "" {
        if id, err := strconv.Atoi(artistaID); err == nil {
            filter.ArtistaID = id
        }
    }

    // Validar estado si está presente
    if filter.Estado != "" {
        validStates := map[string]bool{
            "en venta":  true,
            "vendida":   true,
            "subasta":   true,
            "reservada": true,
        }
        if !validStates[filter.Estado] {
            http.Error(w, "Estado inválido", http.StatusBadRequest)
            return
        }
    }

    report, err := h.repo.GetCategoryArtworks(filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(report)
}