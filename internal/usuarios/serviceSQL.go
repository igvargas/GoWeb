package internal

import "github.com/igvargas/GoWeb/internal/models"

type ServiceSQL interface {
	Store(nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (models.Usuario, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (ssql *serviceSQL) Store(nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (models.Usuario, error) {
	nuevoUsuario := models.Usuario{Nombre: nombre, Apellido: apellido, Email: email, Altura: altura, Activo: activo, FechaCreacion: fecha_creacion}
	usuarioCreado, err := ssql.repository.Store(nuevoUsuario)
	if err != nil {
		return models.Usuario{}, err
	}
	return usuarioCreado, nil
}
