package internal

import (
	"context"

	"github.com/igvargas/GoWeb/internal/models"
)

type ServiceSQL interface {
	Store(nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (models.Usuario, error)
	GetFullData() ([]models.Usuario, error)
	GetOneWithContext(ctx context.Context, id int) (models.Usuario, error)
}

type serviceSQL struct {
	repository RepositorySQL
}

func NewServiceSQL(repo RepositorySQL) ServiceSQL {
	return &serviceSQL{repository: repo}
}

func (ssql *serviceSQL) Store(nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (models.Usuario, error) {
	nuevoUsuario := models.Usuario{Nombre: nombre, Apellido: apellido, Email: email, Edad: edad, Altura: altura, Activo: activo, FechaCreacion: fecha_creacion}
	usuarioCreado, err := ssql.repository.Store(nuevoUsuario)
	if err != nil {
		return models.Usuario{}, err
	}
	return usuarioCreado, nil
}

func (ser *serviceSQL) GetFullData() ([]models.Usuario, error) {
	return ser.repository.GetFullData()
}

func (ser *serviceSQL) GetOneWithContext(ctx context.Context, id int) (models.Usuario, error) {
	return ser.repository.GetOneWithContext(ctx, id)
}
