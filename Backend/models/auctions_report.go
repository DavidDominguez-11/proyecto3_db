// models/auctions_report.go
package models

import "time"

type AuctionOffer struct {
    OfertaID      int       `json:"oferta_id"`
    SubastaID     int       `json:"subasta_id"`
    Obra          string    `json:"obra"`
    Ofertante     string    `json:"ofertante"`
    MontoOfertado float64   `json:"monto_ofertado"`
    FechaOferta   time.Time `json:"fecha_oferta"`
}

type AuctionOffersFilter struct {
    SubastaID   int        `form:"subasta_id"`
    UsuarioID   int        `form:"usuario_id"`
    MontoMin    float64    `form:"monto_min"`
    FechaInicio *time.Time `form:"fecha_inicio"`
}