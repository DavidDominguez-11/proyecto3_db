// repositories/category_repo.go
package repositories

import (
    "fmt"
    "strings"
    "p3db/models"
    "p3db/db"
)

type CategoryRepository struct {
    db *db.Database
}

func NewCategoryRepository(db *db.Database) *CategoryRepository {
    return &CategoryRepository{db: db}
}

func (r *CategoryRepository) GetCategoryArtworks(filter models.CategoryArtworkFilter) ([]models.CategoryArtworkReport, error) {
    baseQuery := `
        SELECT 
            oc.categoria_id,
            c.nombre,
            o.obra_id,
            o.titulo,
            o.precio_referencia,
            o.estado,
            pa.artista_id,
            COUNT(v.venta_id) AS total_ventas
        FROM ObraCategoria oc
        JOIN Categoria c ON oc.categoria_id = c.categoria_id
        JOIN ObraArte o ON oc.obra_id = o.obra_id
        JOIN PerfilArtista pa ON o.artista_id = pa.artista_id
        LEFT JOIN Venta v ON o.obra_id = v.obra_id
    `

    var conditions []string
    var args []interface{}
    argPos := 1

    // Filtros obligatorios
    conditions = append(conditions, fmt.Sprintf("oc.categoria_id = $%d", argPos))
    args = append(args, filter.CategoriaID)
    argPos++

    // Filtros opcionales
    if filter.PrecioMax > 0 {
        conditions = append(conditions, fmt.Sprintf("o.precio_referencia <= $%d", argPos))
        args = append(args, filter.PrecioMax)
        argPos++
    }
    
    if filter.Estado != "" {
        conditions = append(conditions, fmt.Sprintf("o.estado = $%d", argPos))
        args = append(args, filter.Estado)
        argPos++
    }
    
    if filter.ArtistaID > 0 {
        conditions = append(conditions, fmt.Sprintf("pa.artista_id = $%d", argPos))
        args = append(args, filter.ArtistaID)
        argPos++
    }

    // Construir query final
    baseQuery += " WHERE " + strings.Join(conditions, " AND ")
    baseQuery += " GROUP BY oc.categoria_id, c.nombre, o.obra_id, pa.artista_id"
    baseQuery += " ORDER BY total_ventas DESC"

    // Ejecutar consulta
    rows, err := r.db.GetDB().Query(baseQuery, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var report []models.CategoryArtworkReport
    for rows.Next() {
        var item models.CategoryArtworkReport
        err := rows.Scan(
            &item.CategoriaID,
            &item.CategoriaNombre,
            &item.ObraID,
            &item.Titulo,
            &item.PrecioReferencia,
            &item.Estado,
            &item.ArtistaID,
            &item.TotalVentas,
        )
        if err != nil {
            return nil, err
        }
        report = append(report, item)
    }
    
    return report, nil
}