// handlers/sales_handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "p3db/models"
    "p3db/repositories"
    "time"
)

type SalesHandler struct {
    repo *repositories.SaleRepository
}

func NewSalesHandler(repo *repositories.SaleRepository) *SalesHandler {
    return &SalesHandler{repo: repo}
}

func (h *SalesHandler) GetSalesReport(w http.ResponseWriter, r *http.Request) {
    var filter models.SalesReportFilter
    
    // Parsear parámetros
    q := r.URL.Query()
    
    if fechaInicio := q.Get("fecha_inicio"); fechaInicio != "" {
        if t, err := time.Parse(time.RFC3339, fechaInicio); err == nil {
            filter.FechaInicio = &t
        }
    }
    
    filter.MetodoPago = q.Get("metodo_pago")
    filter.PaisArtista = q.Get("pais_artista")
    filter.EstadoEnvio = q.Get("estado_envio")

    // Obtener reporte
    report, err := h.repo.GetSalesReport(filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(report)
}