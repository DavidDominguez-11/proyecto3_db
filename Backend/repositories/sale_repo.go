// repositories/sale_repo.go
package repositories

import (
	"database/sql"
	"fmt"
	"p3db/models"
	"p3db/db"
	"strings"
)

type SaleRepository struct {
    db *db.Database
}

func NewSaleRepository(db *db.Database) *SaleRepository {
    return &SaleRepository{db: db}
}

func (r *SaleRepository) GetSalesReport(filter models.SalesReportFilter) ([]models.SalesReportItem, error) {
    baseQuery := `
        SELECT 
            v.venta_id,
            u.nombre AS comprador,
            o.titulo AS obra,
            pa.pais_origen AS pais_artista,
            v.monto,
            v.metodo_pago,
            v.fecha_venta,
            e.estado_envio
        FROM Venta v
        JOIN Usuario u ON v.usuario_id = u.usuario_id
        JOIN ObraArte o ON v.obra_id = o.obra_id
        JOIN PerfilArtista pa ON o.artista_id = pa.artista_id
        LEFT JOIN Envio e ON v.venta_id = e.venta_id
    `

    var args []interface{}
    conditions := []string{}
    argPos := 1

    // Filtros dinámicos
    if filter.FechaInicio != nil {
        conditions = append(conditions, fmt.Sprintf("v.fecha_venta >= $%d", argPos))
        args = append(args, *filter.FechaInicio)
        argPos++
    }
    
    if filter.MetodoPago != "" {
        conditions = append(conditions, fmt.Sprintf("v.metodo_pago = $%d", argPos))
        args = append(args, filter.MetodoPago)
        argPos++
    }
    
    if filter.PaisArtista != "" {
        conditions = append(conditions, fmt.Sprintf("pa.pais_origen = $%d", argPos))
        args = append(args, filter.PaisArtista)
        argPos++
    }
    
    if filter.EstadoEnvio != "" {
        conditions = append(conditions, fmt.Sprintf("e.estado_envio = $%d", argPos))
        args = append(args, filter.EstadoEnvio)
        argPos++
    }

    // Construir query final
    if len(conditions) > 0 {
        baseQuery += " WHERE " + strings.Join(conditions, " AND ")
    }

    // Ejecutar query
    rows, err := r.db.GetDB().Query(baseQuery, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var report []models.SalesReportItem
    for rows.Next() {
        var item models.SalesReportItem
        var envio sql.NullString
        
        err := rows.Scan(
            &item.VentaID,
            &item.Comprador,
            &item.Obra,
            &item.PaisArtista,
            &item.Monto,
            &item.MetodoPago,
            &item.FechaVenta,
            &envio,
        )
        
        if err != nil {
            return nil, err
        }
        
        item.EstadoEnvio = envio.String
        report = append(report, item)
    }
    
    return report, nil
}