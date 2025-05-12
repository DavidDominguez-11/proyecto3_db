// repositories/transaction_repo.go
package repositories

import (
    "fmt"
    "strings"
    "p3db/models"
    "p3db/db"
)

type TransactionRepository struct {
    db *db.Database
}

func NewTransactionRepository(db *db.Database) *TransactionRepository {
    return &TransactionRepository{db: db}
}

func (r *TransactionRepository) GetTransactions(filter models.TransactionFilter) ([]models.Transaction, error) {
    baseQuery := `
        SELECT 
            transaccion_id,
            tipo,
            fecha,
            entidad_afectada_id,
            detalle
        FROM Transaccion
    `

    var conditions []string
    var args []interface{}
    argPos := 1

    // Construir condiciones dinámicas
    if filter.Tipo != "" {
        conditions = append(conditions, fmt.Sprintf("tipo = $%d", argPos))
        args = append(args, filter.Tipo)
        argPos++
    }
    
    if filter.FechaInicio != nil {
        conditions = append(conditions, fmt.Sprintf("fecha >= $%d", argPos))
        args = append(args, filter.FechaInicio)
        argPos++
    }
    
    if filter.TransaccionID > 0 {
        conditions = append(conditions, fmt.Sprintf("transaccion_id = $%d", argPos))
        args = append(args, filter.TransaccionID)
        argPos++
    }

    if filter.EntidadAfectadaID > 0 {
        conditions = append(conditions, fmt.Sprintf("entidad_afectada_id = $%d", argPos))
        args = append(args, filter.EntidadAfectadaID)
        argPos++
    }

    // Combinar condiciones
    if len(conditions) > 0 {
        baseQuery += " WHERE " + strings.Join(conditions, " AND ")
    }

    // Ordenar por fecha descendente
    baseQuery += " ORDER BY fecha DESC"

    // Ejecutar consulta
    rows, err := r.db.GetDB().Query(baseQuery, args...)
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var transactions []models.Transaction
    for rows.Next() {
        var t models.Transaction
        err := rows.Scan(
            &t.ID,
            &t.Tipo,
            &t.Fecha,
            &t.EntidadAfectadaID,
            &t.Detalle,
        )
        if err != nil {
            return nil, err
        }
        transactions = append(transactions, t)
    }
    
    return transactions, nil
}