// repositories/artist_repo.go
package repositories

import (
	"p3db/db"
	"p3db/models"
)

// ArtistRepository maneja operaciones CRUD para PerfilArtista
type ArtistRepository struct {
	db *db.Database
}

// NewArtistRepository crea un nuevo repositorio de artistas
func NewArtistRepository(db *db.Database) *ArtistRepository {
	return &ArtistRepository{db: db}
}

// Create inserta un nuevo perfil de artista
func (r *ArtistRepository) Create(a *models.PerfilArtista) error {
	query := `INSERT INTO PerfilArtista (usuario_id, biografia, pais_origen, estilo_principal)
	          VALUES ($1, $2, $3, $4)
	          RETURNING artista_id` 
	return r.db.GetDB().QueryRow(query, a.UsuarioID, a.Biografia, a.PaisOrigen, a.EstiloPrincipal).
		Scan(&a.ID)
}

// GetAll obtiene todos los perfiles de artista
func (r *ArtistRepository) GetAll() ([]models.PerfilArtista, error) {
	rows, err := r.db.GetDB().Query(
		"SELECT artista_id, usuario_id, biografia, pais_origen, estilo_principal FROM PerfilArtista",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []models.PerfilArtista
	for rows.Next() {
		var a models.PerfilArtista
		err := rows.Scan(&a.ID, &a.UsuarioID, &a.Biografia, &a.PaisOrigen, &a.EstiloPrincipal)
		if err != nil {
			return nil, err
		}
		list = append(list, a)
	}
	return list, nil
}

// GetByID obtiene un perfil por su ID
func (r *ArtistRepository) GetByID(id int) (*models.PerfilArtista, error) {
	a := &models.PerfilArtista{}
	query := `SELECT artista_id, usuario_id, biografia, pais_origen, estilo_principal
	          FROM PerfilArtista WHERE artista_id = $1`
	err := r.db.GetDB().QueryRow(query, id).
		Scan(&a.ID, &a.UsuarioID, &a.Biografia, &a.PaisOrigen, &a.EstiloPrincipal)
	if err != nil {
		return nil, err
	}
	return a, nil
}

// Update modifica un perfil existente
func (r *ArtistRepository) Update(a *models.PerfilArtista) error {
	query := `UPDATE PerfilArtista SET biografia = $1, pais_origen = $2, estilo_principal = $3
	          WHERE artista_id = $4`
	_, err := r.db.GetDB().Exec(query, a.Biografia, a.PaisOrigen, a.EstiloPrincipal, a.ID)
	return err
}

// Delete elimina un perfil de artista
func (r *ArtistRepository) Delete(id int) error {
	_, err := r.db.GetDB().Exec("DELETE FROM PerfilArtista WHERE artista_id = $1", id)
	return err
}