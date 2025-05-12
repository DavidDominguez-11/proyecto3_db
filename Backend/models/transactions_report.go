// models/transactions_report.go
package models

import "time"

type Transaction struct {
    ID                int            `json:"transaccion_id"`
    Tipo              string         `json:"tipo"`
    Fecha             time.Time      `json:"fecha"`
    EntidadAfectadaID int            `json:"entidad_afectada_id"`
    Detalle           string         `json:"detalle,omitempty"`
}

type TransactionFilter struct {
    TransaccionID      int        `form:"transaccion_id"`
    Tipo               string     `form:"tipo"`
    FechaInicio        *time.Time `form:"fecha_inicio"`
    EntidadAfectadaID  int        `form:"entidad_afectada_id"`
}