// handlers/auction_handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "time"
    "p3db/models"
    "p3db/repositories"
    
    "github.com/gorilla/mux"
)

type AuctionHandler struct {
    repo *repositories.AuctionRepository
}

func NewAuctionHandler(repo *repositories.AuctionRepository) *AuctionHandler {
    return &AuctionHandler{repo: repo}
}

func (h *AuctionHandler) GetAuctionOffers(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    subastaID, _ := strconv.Atoi(vars["subasta_id"])
    
    var filter models.AuctionOffersFilter
    filter.SubastaID = subastaID
    
    q := r.URL.Query()
    
    // Parsear parámetros
    if usuarioID := q.Get("usuario_id"); usuarioID != "" {
        if id, err := strconv.Atoi(usuarioID); err == nil {
            filter.UsuarioID = id
        }
    }
    
    if montoMin := q.Get("monto_min"); montoMin != "" {
        if monto, err := strconv.ParseFloat(montoMin, 64); err == nil {
            filter.MontoMin = monto
        }
    }
    
    if fechaInicio := q.Get("fecha_inicio"); fechaInicio != "" {
        if t, err := time.Parse(time.RFC3339, fechaInicio); err == nil {
            filter.FechaInicio = &t
        }
    }

    offers, err := h.repo.GetAuctionOffers(filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(offers)
}