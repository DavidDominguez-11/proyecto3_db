// models/category_report.go
package models

type CategoryArtworkReport struct {
    CategoriaID     int     `json:"categoria_id"`
    CategoriaNombre string  `json:"categoria"`
    ObraID          int     `json:"obra_id"`
    Titulo          string  `json:"titulo"`
    PrecioReferencia float64 `json:"precio_referencia"`
    Estado          string  `json:"estado"`
    ArtistaID       int     `json:"artista_id"`
    TotalVentas     int     `json:"total_ventas"`
}

type CategoryArtworkFilter struct {
    CategoriaID int     `form:"categoria_id"`
    PrecioMax   float64 `form:"precio_max"`
    Estado      string  `form:"estado"`
    ArtistaID   int     `form:"artista_id"`
}