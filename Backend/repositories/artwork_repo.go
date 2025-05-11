package repositories

import (
	"p3db/db"
	"p3db/models"
)

// ArtworkRepository maneja operaciones CRUD para ObraArte
type ArtworkRepository struct {
	db *db.Database
}

// NewArtworkRepository crea un nuevo repositorio de arte
func NewArtworkRepository(db *db.Database) *ArtworkRepository {
	return &ArtworkRepository{db: db}
}

// Create inserta una nueva obra de arte
func (r *ArtworkRepository) Create(a *models.ObraArte) error {
	query := `INSERT INTO ObraArte (titulo, descripcion, año_creacion, precio_referencia, estado, artista_id) 
	          VALUES ($1, $2, $3, $4, $5, $6) RETURNING obra_id`
	return r.db.GetDB().QueryRow(query,
		a.Titulo,
		a.Descripcion,
		a.AñoCreacion,
		a.PrecioReferencia,
		a.Estado,
		a.ArtistaID,
	).Scan(&a.ID)
}

// GetAll obtiene todas las obras de arte
func (r *ArtworkRepository) GetAll() ([]models.ObraArte, error) {
	rows, err := r.db.GetDB().Query(
		"SELECT obra_id, titulo, descripcion, año_creacion, precio_referencia, estado, artista_id FROM ObraArte",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.ObraArte
	for rows.Next() {
		var a models.ObraArte
		err := rows.Scan(&a.ID, &a.Titulo, &a.Descripcion, &a.AñoCreacion,
			&a.PrecioReferencia, &a.Estado, &a.ArtistaID)
		if err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

// GetByID obtiene una obra por su ID
func (r *ArtworkRepository) GetByID(id int) (*models.ObraArte, error) {
	a := &models.ObraArte{}
	query := `SELECT obra_id, titulo, descripcion, año_creacion, precio_referencia, estado, artista_id
	          FROM ObraArte WHERE obra_id = $1`
	err := r.db.GetDB().QueryRow(query, id).
		Scan(&a.ID, &a.Titulo, &a.Descripcion, &a.AñoCreacion,
			 &a.PrecioReferencia, &a.Estado, &a.ArtistaID)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Update modifica una obra existente
func (r *ArtworkRepository) Update(a *models.ObraArte) error {
	query := `UPDATE ObraArte SET titulo = $1, descripcion = $2, año_creacion = $3,
		 precio_referencia = $4, estado = $5, artista_id = $6 WHERE obra_id = $7`
	_, err := r.db.GetDB().Exec(query,
		a.Titulo, a.Descripcion, a.AñoCreacion, a.PrecioReferencia, a.Estado, a.ArtistaID, a.ID)
	return err
}

// Delete elimina una obra de arte
func (r *ArtworkRepository) Delete(id int) error {
	_, err := r.db.GetDB().Exec("DELETE FROM ObraArte WHERE obra_id = $1", id)
	return err
}
