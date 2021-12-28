package internal

import (
	"context"
	"database/sql"
	"log"

	"github.com/igvargas/GoWeb/internal/models"
	"github.com/igvargas/GoWeb/pkg/db"
)

type RepositorySQL interface {
	Store(usuario models.Usuario) (models.Usuario, error)
	GetFullData() ([]models.Usuario, error)
	GetOneWithContext(ctx context.Context, id int) (models.Usuario, error)
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

func (r *repositorySQL) GetFullData() ([]models.Usuario, error) {
	var misUsuarios []models.Usuario
	db := db.StorageDB
	var usuarioLeido models.Usuario
	rows, err := db.Query("SELECT u.nombre, u.apellido, u.edad, c.nombre_ciudad, c.nombre_pais FROM usuarios u inner join ciudades c on u.idciudad = c.id;")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(&usuarioLeido.Nombre, &usuarioLeido.Apellido, &usuarioLeido.Edad, &usuarioLeido.Domicilio.NombreCiudad, &usuarioLeido.Domicilio.NombrePais)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		misUsuarios = append(misUsuarios, usuarioLeido)
	}
	return misUsuarios, nil
}

func (r *repositorySQL) GetOneWithContext(ctx context.Context, id int) (models.Usuario, error) {
	db := db.StorageDB
	var usuarioLeido models.Usuario
	rows, err := db.QueryContext(ctx, "SELECT id, nombre,apellido, edad FROM usuarios WHERE id = ?", id)

	if err != nil {
		log.Fatal(err)
		return usuarioLeido, err
	}

	for rows.Next() {
		err = rows.Scan(&usuarioLeido.ID, &usuarioLeido.Nombre, &usuarioLeido.Apellido, &usuarioLeido.Edad)
		if err != nil {
			log.Fatal(err)
			return usuarioLeido, err
		}

	}
	return usuarioLeido, nil
}
