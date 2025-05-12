package models

// models/artworks_report.go
type ArtworkReportItem struct {
    ObraID          int       `json:"obra_id"`
    Titulo          string    `json:"titulo"`
    Estado          string    `json:"estado"`
    PrecioReferencia float64   `json:"precio_referencia"`
    AñoCreacion     int       `json:"año_creacion"`
    EstiloPrincipal string    `json:"estilo_principal"`
    PaisOrigen      string    `json:"pais_origen"`
}

type ArtworkFilter struct {
    Estado          string   `form:"estado"`
    EstiloPrincipal string   `form:"estilo"`
    PaisOrigen      string   `form:"pais"`
    PrecioMin       float64  `form:"precio_min"`
    PrecioMax       float64  `form:"precio_max"`
}