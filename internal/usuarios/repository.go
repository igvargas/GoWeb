package internal

import (
	"fmt"

	"github.com/igvargas/GoWeb/pkg/store"
)

type Usuario struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int     `json:"edad"`
	Altura        float64 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fecha_creacion"`
}

var usuarios []Usuario

type Repository interface {
	GetAll() ([]Usuario, error)
	Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error)
	LastId() (int, error)
	Update(id int, nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error)
	UpdateNombre(id int, nombre string) (Usuario, error)
	Delete(id int) error
}

type repository struct {
	db store.Store
}

func NewRepository(db store.Store) Repository {
	return &repository{db}
}

func (repo *repository) GetAll() ([]Usuario, error) {
	err := repo.db.Read(&usuarios)
	if err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (repo *repository) Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error) {
	repo.db.Read(&usuarios)

	usr := Usuario{id, nombre, apellido, email, edad, altura, activo, fecha_creacion}

	usuarios = append(usuarios, usr)
	err := repo.db.Write(usuarios)

	if err != nil {
		return Usuario{}, err
	}

	return usr, nil
}

func (repo *repository) LastId() (int, error) {
	err := repo.db.Read(&usuarios)
	if err != nil {
		return 0, err
	}

	if len(usuarios) == 0 {
		return 0, nil
	}

	lastID := usuarios[len(usuarios)-1].ID
	return lastID, nil
}

func (repo *repository) Update(id int, nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error) {
	err := repo.db.Read(&usuarios)
	if err != nil {
		return Usuario{}, err
	}

	usr := Usuario{id, nombre, apellido, email, edad, altura, activo, fecha_creacion}
	for i, v := range usuarios {
		if v.ID == id {
			usuarios[i] = usr
			err := repo.db.Write(usuarios)

			if err != nil {
				return Usuario{}, err
			}
			return usr, nil
		}
	}

	return Usuario{}, fmt.Errorf("El usuario %d no existe", id)
}

func (repo *repository) UpdateNombre(id int, nombre string) (Usuario, error) {
	err := repo.db.Read(&usuarios)
	if err != nil {
		return Usuario{}, err
	}
	for i, v := range usuarios {
		if v.ID == id {
			usuarios[i].Nombre = nombre
			err := repo.db.Write(usuarios)

			if err != nil {
				return Usuario{}, err
			}
			return usuarios[i], nil
		}
	}
	return Usuario{}, fmt.Errorf("El usuario %d no existe", id)

}

func (repo *repository) Delete(id int) error {
	err := repo.db.Read(&usuarios)
	if err != nil {
		return err
	}
	index := 0
	for i, v := range usuarios {
		if v.ID == id {
			index = i
			usuarios = append(usuarios[:index], usuarios[index+1:]...)
			err := repo.db.Write(usuarios)
			return err
		}
	}
	return fmt.Errorf("La persona %d no existe", id)
}
