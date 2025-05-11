// models/sales_report.go
package models

import "time"

type SalesReportItem struct {
    VentaID       int       `json:"venta_id"`
    Comprador     string    `json:"comprador"`
    Obra          string    `json:"obra"`
    PaisArtista   string    `json:"pais_artista"`
    Monto         float64   `json:"monto"`
    MetodoPago    string    `json:"metodo_pago"`
    FechaVenta    time.Time `json:"fecha_venta"`
    EstadoEnvio   string    `json:"estado_envio"`
}

type SalesReportFilter struct {
    FechaInicio   *time.Time `form:"fecha_inicio"`
    FechaFin      *time.Time `form:"fecha_fin"`
    MetodoPago    string     `form:"metodo_pago"`
    PaisArtista   string     `form:"pais_artista"`
    EstadoEnvio   string     `form:"estado_envio"`
}