package internal

import (
	"database/sql"
	"log"

	"github.com/igvargas/GoWeb/internal/models"
	"github.com/igvargas/GoWeb/pkg/db"
)

type RepositorySQL interface {
	Store(usuario models.Usuario) (models.Usuario, error)
}

type repositorySQL struct{}

func NewRepositorySQL() RepositorySQL {
	return &repositorySQL{}
}

func (r *repositorySQL) Store(usuario models.Usuario) (models.Usuario, error) {
	db := db.StorageDB

	stmt, err := db.Prepare("INSERT INTO usuarios(nombre, apellido, email, edad, altura, activo, fecha_creacion) VALUES (?,?,?,?,?,?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var result sql.Result
	result, err = stmt.Exec(usuario.Nombre, usuario.Apellido, usuario.Email, usuario.Edad, usuario.Altura, usuario.Activo, usuario.FechaCreacion)
	if err != nil {
		return models.Usuario{}, err
	}
	idCreado, _ := result.LastInsertId()
	usuario.ID = int(idCreado)

	return usuario, nil
}
