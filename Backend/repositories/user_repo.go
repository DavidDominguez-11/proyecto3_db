// repositories/user_repo.go
package repositories

import (
	"p3db/models"
	"p3db/db"
)

type UserRepository struct {
	db *db.Database
}

func NewUserRepository(db *db.Database) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Create(user *models.Usuario) error {
	query := `INSERT INTO Usuario (nombre, correo, tipo_usuario) 
	          VALUES ($1, $2, $3) RETURNING usuario_id, fecha_registro`
	
	err := r.db.GetDB().QueryRow(query, user.Nombre, user.Correo, user.TipoUsuario).
		Scan(&user.ID, &user.FechaRegistro)
	return err
}

func (r *UserRepository) GetAll() ([]models.Usuario, error) {
	rows, err := r.db.GetDB().Query("SELECT usuario_id, nombre, correo, fecha_registro, tipo_usuario FROM Usuario")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.Usuario
	for rows.Next() {
		var user models.Usuario
		err := rows.Scan(&user.ID, &user.Nombre, &user.Correo, &user.FechaRegistro, &user.TipoUsuario)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *UserRepository) GetByID(id int) (*models.Usuario, error) {
	user := &models.Usuario{}
	query := `SELECT usuario_id, nombre, correo, fecha_registro, tipo_usuario 
	          FROM Usuario WHERE usuario_id = $1`
	
	err := r.db.GetDB().QueryRow(query, id).Scan(
		&user.ID, &user.Nombre, &user.Correo, &user.FechaRegistro, &user.TipoUsuario)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserRepository) Update(user *models.Usuario) error {
	query := `UPDATE Usuario SET nombre = $1, correo = $2, tipo_usuario = $3 
	          WHERE usuario_id = $4`
	_, err := r.db.GetDB().Exec(query, user.Nombre, user.Correo, user.TipoUsuario, user.ID)
	return err
}

func (r *UserRepository) Delete(id int) error {
	_, err := r.db.GetDB().Exec("DELETE FROM Usuario WHERE usuario_id = $1", id)
	return err
}