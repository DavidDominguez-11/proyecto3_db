// repositories/auction_repo.go
package repositories

import (
    "fmt"
    "strings"
    "p3db/models"
    "p3db/db"
)

type AuctionRepository struct {
    db *db.Database
}

func NewAuctionRepository(db *db.Database) *AuctionRepository {
    return &AuctionRepository{db: db}
}

func (r *AuctionRepository) GetAuctionOffers(filter models.AuctionOffersFilter) ([]models.AuctionOffer, error) {
    baseQuery := `
        SELECT 
            os.oferta_id,
            s.subasta_id,
            o.titulo AS obra,
            u.nombre AS ofertante,
            os.monto_ofertado,
            os.fecha_oferta
        FROM OfertaSubasta os
        JOIN Subasta s ON os.subasta_id = s.subasta_id
        JOIN ObraArte o ON s.obra_id = o.obra_id
        JOIN Usuario u ON os.usuario_id = u.usuario_id
        WHERE os.subasta_id = $1
    `

    args := []interface{}{filter.SubastaID}
    conditions := []string{}
    argPos := 2

    // Filtros dinámicos
    if filter.FechaInicio != nil {
        conditions = append(conditions, fmt.Sprintf("os.fecha_oferta >= $%d", argPos))
        args = append(args, filter.FechaInicio)
        argPos++
    }
    
    if filter.MontoMin > 0 {
        conditions = append(conditions, fmt.Sprintf("os.monto_ofertado >= $%d", argPos))
        args = append(args, filter.MontoMin)
        argPos++
    }
    
    if filter.UsuarioID > 0 {
        conditions = append(conditions, fmt.Sprintf("os.usuario_id = $%d", argPos))
        args = append(args, filter.UsuarioID)
        argPos++
    }

    // Construir query final
    if len(conditions) > 0 {
        baseQuery += " AND " + strings.Join(conditions, " AND ")
    }

    // Ordenar por monto descendente
    baseQuery += " ORDER BY os.monto_ofertado DESC"

    // Ejecutar query
    rows, err := r.db.GetDB().Query(baseQuery, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var offers []models.AuctionOffer
    for rows.Next() {
        var offer models.AuctionOffer
        err := rows.Scan(
            &offer.OfertaID,
            &offer.SubastaID,
            &offer.Obra,
            &offer.Ofertante,
            &offer.MontoOfertado,
            &offer.FechaOferta,
        )
        
        if err != nil {
            return nil, err
        }
        
        offers = append(offers, offer)
    }
    
    return offers, nil
}