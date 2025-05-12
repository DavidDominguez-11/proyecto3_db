// handlers/transaction_handler.go
package handlers

import (
    "encoding/json"
    "net/http"
    "strconv"
    "time"
    "p3db/models"
    "p3db/repositories"
)

type TransactionHandler struct {
    repo *repositories.TransactionRepository
}

func NewTransactionHandler(repo *repositories.TransactionRepository) *TransactionHandler {
    return &TransactionHandler{repo: repo}
}

func (h *TransactionHandler) GetTransactions(w http.ResponseWriter, r *http.Request) {
    var filter models.TransactionFilter
    q := r.URL.Query()
    
    // Parsear parámetros
    filter.Tipo = q.Get("tipo")
    
    if fechaInicio := q.Get("fecha_inicio"); fechaInicio != "" {
        if t, err := time.Parse(time.RFC3339, fechaInicio); err == nil {
            filter.FechaInicio = &t
        }
    }
    
    if transaccionID := q.Get("transaccion_id"); transaccionID != "" {
        if id, err := strconv.Atoi(transaccionID); err == nil {
            filter.TransaccionID = id
        } else {
            http.Error(w, "transaccion_id inválido", http.StatusBadRequest)
            return
    }
}

    if entidadAfectadaID := q.Get("entidad_afectada_id"); entidadAfectadaID != "" {
        if id, err := strconv.Atoi(entidadAfectadaID); err == nil {
            filter.EntidadAfectadaID = id
        }
    }

    // Validar tipo de transacción
    if filter.Tipo != "" {
        validTypes := map[string]bool{
            "venta":          true,
            "subasta_ganada": true,
            "oferta":         true,
            "cambio_estado":  true,
        }
        if !validTypes[filter.Tipo] {
            http.Error(w, "Tipo de transacción inválido", http.StatusBadRequest)
            return
        }
    }

    transactions, err := h.repo.GetTransactions(filter)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(transactions)
}