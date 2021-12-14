package internal

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/igvargas/GoWeb/pkg/store"
	"github.com/stretchr/testify/assert"
)

var user string = `[
	{"id": 1,"nombre": "josefa","apellido": "Maylor","email": "lmaylor0@dagondesign.com","edad": 30,"altura": 177.1,"activo": false,"fecha_creacion": "8/20/2021"},
    {"id": 2,"nombre": "Lenard","apellido": "Maylor","email": "lmaylor0@dagondesign.com","edad": 30,"altura": 177.1,"activo": false,"fecha_creacion": "8/20/2021"}]`

func TestGetAllServiceMock(t *testing.T) {
	// Arrange
	dataByte := []byte(user)
	var userExpected []Usuario
	json.Unmarshal(dataByte, &userExpected)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	// Act
	misUser, _ := service.GetAll()

	// Assert
	assert.Equal(t, userExpected, misUser)
}

func TestGetAllServiceMockError(t *testing.T) {
	// Arrange
	errorExpected := errors.New("No hay datos en el Mock")

	dbMock := store.Mock{Err: errorExpected}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	// Act
	misUser, errorRecibido := service.GetAll()

	// Assert
	assert.Equal(t, errorExpected, errorRecibido)
	assert.Nil(t, misUser)
}

func TestStoreServiceMock(t *testing.T) {
	// Arrange
	newUser := Usuario{
		Nombre:        "cristian",
		Apellido:      "lopez",
		Email:         "clopez@hotmail.com",
		Edad:          41,
		Altura:        1.78,
		Activo:        true,
		FechaCreacion: "17/05/2001",
	}

	dbMock := store.Mock{Data: []byte(`[]`)}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	// Act
	userCreado, _ := service.Store(newUser.Nombre, newUser.Apellido, newUser.Email, newUser.Edad, newUser.Altura, newUser.Activo, newUser.FechaCreacion)

	// Assert
	assert.Equal(t, newUser.Nombre, userCreado.Nombre)
}

func TestStoreServiceMockError(t *testing.T) {
	// Arrange
	errorExpected := errors.New("No hay datos en el Mock")
	newUser := Usuario{
		Nombre:        "cristian",
		Apellido:      "lopez",
		Email:         "clopez@hotmail.com",
		Edad:          41,
		Altura:        1.78,
		Activo:        true,
		FechaCreacion: "17/05/2001",
	}

	dbMock := store.Mock{Data: []byte(`[]`), Err: errorExpected}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	// Act
	userCreado, err := service.Store(newUser.Nombre, newUser.Apellido, newUser.Email, newUser.Edad, newUser.Altura, newUser.Activo, newUser.FechaCreacion)

	// Assert
	assert.Equal(t, errorExpected, err)
	assert.Equal(t, Usuario{}, userCreado)
}

func TestServiceUpdate(t *testing.T) {
	// Arrange
	var user string = `[{"id":1,"nombre":"adrian","apellido":"lopez","email":"clopez@hotmail.com","edad":41,"altura":1.78,"activo":true,"fecha_creacion":"17/05/2001"}]`

	patchUser := Usuario{
		Nombre:        "cristian",
		Apellido:      "lopez",
		Email:         "clopez@hotmail.com",
		Edad:          41,
		Altura:        1.78,
		Activo:        true,
		FechaCreacion: "17/05/2001",
	}
	dataByte := []byte(user)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	service := NewService(repo)

	// Act
	userActualizado, _ := service.Update(1, patchUser.Nombre, patchUser.Apellido, patchUser.Email, patchUser.Edad, patchUser.Altura, patchUser.Activo, patchUser.FechaCreacion)

	userExpected := Usuario{
		ID:            1,
		Nombre:        "cristian",
		Apellido:      "lopez",
		Email:         "clopez@hotmail.com",
		Edad:          41,
		Altura:        1.78,
		Activo:        true,
		FechaCreacion: "17/05/2001",
	}

	// Assert
	assert.Equal(t, userExpected, userActualizado)
}

func TestServiceDelete(t *testing.T) {
	// Arrange
	dataByte := []byte(user)
	var userExpected []Usuario
	json.Unmarshal(dataByte, &userExpected)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	// Act
	err := service.Delete(1)

	// Assert
	assert.Nil(t, err)
}

func TestServiceDeleteError(t *testing.T) {
	// Arrange
	dataByte := []byte(user)
	var userExpected []Usuario
	json.Unmarshal(dataByte, &userExpected)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)
	service := NewService(repo)

	// Act
	err := service.Delete(12)

	// Assert
	assert.Error(t, err)
}
