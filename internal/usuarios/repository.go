package internal

import "fmt"

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
var lastID int

type Repository interface {
	GetAll() ([]Usuario, error)
	Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error)
	LastId() (int, error)
	Update(id int, nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error)
	UpdateNombre(id int, nombre string) (Usuario, error)
	Delete(id int) error
}

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (repo *repository) GetAll() ([]Usuario, error) {
	return usuarios, nil
}

func (repo *repository) Store(id int, nombre string, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error) {
	usr := Usuario{id, nombre, apellido, email, edad, altura, activo, fecha_creacion}
	lastID = id
	usuarios = append(usuarios, usr)
	return usr, nil
}

func (repo *repository) LastId() (int, error) {
	return lastID, nil
}

func (repo *repository) Update(id int, nombre, apellido string, email string, edad int, altura float64, activo bool, fecha_creacion string) (Usuario, error) {
	usr := Usuario{id, nombre, apellido, email, edad, altura, activo, fecha_creacion}
	for i, v := range usuarios {
		if v.ID == id {
			usuarios[i] = usr
			return usr, nil
		}
	}
	return Usuario{}, fmt.Errorf("El usuario %d no existe", id)
}

func (repo *repository) UpdateNombre(id int, nombre string) (Usuario, error) {
	for i, v := range usuarios {
		if v.ID == id {
			usuarios[i].Nombre = nombre
			return usuarios[i], nil
		}
	}
	return Usuario{}, fmt.Errorf("El usuario %d no existe", id)

}

func (repo *repository) Delete(id int) error {
	index := 0
	for i, v := range usuarios {
		if v.ID == id {
			index = i
			usuarios = append(usuarios[:index], usuarios[index+1:]...)
			return nil
		}
	}
	return fmt.Errorf("La persona %d no existe", id)
}
