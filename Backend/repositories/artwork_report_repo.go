package repositories

import (
	"fmt"
	"p3db/models"
	"p3db/db"
	"strings"
)

type AtworkReportRepository struct {
    db *db.Database
}

func NewAtworkReportRepository(db *db.Database) *AtworkReportRepository {
    return &AtworkReportRepository{db: db}
}

// repositories/AtworkReportRepository_repo.go
func (r *ArtworkRepository) GetFilteredArtworks(filter models.ArtworkFilter) ([]models.ArtworkReportItem, error) {
    baseQuery := `
        SELECT 
            o.obra_id,
            o.titulo,
            o.estado,
            o.precio_referencia,
            o.año_creacion,
            pa.estilo_principal,
            pa.pais_origen
        FROM ObraArte o
        JOIN PerfilArtista pa ON o.artista_id = pa.artista_id
    `

    var conditions []string
    var args []interface{}
    argPos := 1

    // Filtros dinámicos
    if filter.Estado != "" {
        conditions = append(conditions, fmt.Sprintf("o.estado = $%d", argPos))
        args = append(args, filter.Estado)
        argPos++
    }
    
    if filter.EstiloPrincipal != "" {
        conditions = append(conditions, fmt.Sprintf("pa.estilo_principal = $%d", argPos))
        args = append(args, filter.EstiloPrincipal)
        argPos++
    }
    
    if filter.PaisOrigen != "" {
        conditions = append(conditions, fmt.Sprintf("pa.pais_origen = $%d", argPos))
        args = append(args, filter.PaisOrigen)
        argPos++
    }
    
    if filter.PrecioMin > 0 || filter.PrecioMax > 0 {
        if filter.PrecioMin > 0 && filter.PrecioMax > 0 {
            conditions = append(conditions, fmt.Sprintf("o.precio_referencia BETWEEN $%d AND $%d", argPos, argPos+1))
            args = append(args, filter.PrecioMin, filter.PrecioMax)
            argPos += 2
        } else if filter.PrecioMin > 0 {
            conditions = append(conditions, fmt.Sprintf("o.precio_referencia >= $%d", argPos))
            args = append(args, filter.PrecioMin)
            argPos++
        } else {
            conditions = append(conditions, fmt.Sprintf("o.precio_referencia <= $%d", argPos))
            args = append(args, filter.PrecioMax)
            argPos++
        }
    }

    if len(conditions) > 0 {
        baseQuery += " WHERE " + strings.Join(conditions, " AND ")
    }

    // Ordenar por precio descendente
    baseQuery += " ORDER BY o.precio_referencia DESC"

    rows, err := r.db.GetDB().Query(baseQuery, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var artworks []models.ArtworkReportItem
    for rows.Next() {
        var a models.ArtworkReportItem
        err := rows.Scan(
            &a.ObraID,
            &a.Titulo,
            &a.Estado,
            &a.PrecioReferencia,
            &a.AñoCreacion,
            &a.EstiloPrincipal,
            &a.PaisOrigen,
        )
        if err != nil {
            return nil, err
        }
        artworks = append(artworks, a)
    }
    
    return artworks, nil
}