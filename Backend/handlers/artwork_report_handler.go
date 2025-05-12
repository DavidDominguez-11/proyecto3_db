package handlers

import (
    "encoding/json"
    "net/http"
    "p3db/models"
	"strconv"

)

// handlers/artwork_report_handler.go
func (h *ArtworkHandler) GetArtworkReport(w http.ResponseWriter, r *http.Request) {
    var filter models.ArtworkFilter
    
    q := r.URL.Query()
    
    // Parsear parámetros
    filter.Estado = q.Get("estado")
    filter.EstiloPrincipal = q.Get("estilo")
    filter.PaisOrigen = q.Get("pais")
    
    if precioMin := q.Get("precio_min"); precioMin != "" {
        if min, err := strconv.ParseFloat(precioMin, 64); err == nil {
            filter.PrecioMin = min
        }
    }
    
    if precioMax := q.Get("precio_max"); precioMax != "" {
        if max, err := strconv.ParseFloat(precioMax, 64); err == nil {
            filter.PrecioMax = max
        }
    }

    // Validar estado si está presente
    if filter.Estado != "" {
        validStates := map[string]bool{
            "en venta": true,
            "subasta":  true,
            "vendida":  true,
            "reservada": true,
        }
        if !validStates[filter.Estado] {
            http.Error(w, "Estado inválido", http.StatusBadRequest)
            return
        }
    }

    artworks, err := h.repo.GetFilteredArtworks(filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(artworks)
}