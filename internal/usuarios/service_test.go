package internal

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/igvargas/GoWeb/internal/models"
	"github.com/igvargas/GoWeb/pkg/db"
	"github.com/igvargas/GoWeb/pkg/store"
	"github.com/stretchr/testify/assert"
)

var user string = `[
	{"id": 1,"nombre": "josefa","apellido": "Maylor","email": "lmaylor0@dagondesign.com","edad": 30,"altura": 177.1,"activo": false,"fecha_creacion": "8/20/2021"},
    {"id": 2,"nombre": "Lenard","apellido": "Maylor","email": "lmaylor0@dagondesign.com","edad": 30,"altura": 177.1,"activo": false,"fecha_creacion": "8/20/2021"}]`

func TestGetAllServiceMock(t *testing.T) {
	// Arrange
	// dataByte := []byte(user)
	// var userExpected []Usuario
	// json.Unmarshal(dataByte, &userExpected)
	sliceDeBytes, _ := json.Marshal(sliceDeUsuarios)

	dbMock := store.Mock{Data: sliceDeBytes}
	storeStub := store.FileStore{FileName: "", Mock: &dbMock}

	repo := NewRepository(&storeStub)
	service := NewService(repo)

	// Act
	resul, err := service.GetAll()

	// Assert
	assert.Nil(t, err)
	assert.True(t, len(resul) == 2)
}

func TestGetAllServiceMockError(t *testing.T) {
	// Arrange
	errorExpected := errors.New("No hay datos en el Mock")

	dbMock := store.Mock{Err: errorExpected}
	storeStub := store.FileStore{FileName: "", Mock: &dbMock}

	repo := NewRepository(&storeStub)
	service := NewService(repo)

	// Act
	_, errorRecibido := service.GetAll()

	// Assert
	assert.Error(t, errorRecibido)
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

func TestStoreServiceSQL(t *testing.T) {
	newUser := models.Usuario{
		Nombre:        "jose",
		Apellido:      "lopez",
		Email:         "jlopez@hotmail.com",
		Edad:          41,
		Altura:        1.68,
		Activo:        true,
		FechaCreacion: "8/05/1966",
	}

	repo := NewRepositorySQL()
	service := NewServiceSQL(repo)

	usuarioCreado, _ := service.Store(newUser.Nombre, newUser.Apellido, newUser.Email, newUser.Edad, newUser.Altura, newUser.Activo, newUser.FechaCreacion)

	assert.Equal(t, newUser.Nombre, usuarioCreado.Nombre)
}

func TestGetFullDataServiceSQL(t *testing.T) {
	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	misUsuarios, err := service.GetFullData()
	assert.Nil(t, err)
	assert.Equal(t, "cordoba", misUsuarios[0].Domicilio.NombreCiudad)
	assert.True(t, len(misUsuarios) > 0)
}

func TestGetOneContextServiceSQL(t *testing.T) {
	usuarioNuevo := models.Usuario{
		Nombre: "jose",
	}

	repo := NewRepositorySQL()

	service := NewServiceSQL(repo)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usuarioCargado, err := service.GetOneWithContext(ctx, 2)

	assert.Nil(t, err)
	assert.Equal(t, usuarioNuevo.Nombre, usuarioCargado.Nombre)
}

func TestGetOneContextServiceSQLMock(t *testing.T) {
	usuarioNuevo := models.Usuario{
		Nombre: "jose",
	}
	db := db.StorageDB
	repo := NewRepositorySQLMock(db)

	service := NewServiceSQL(repo)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	usuarioCargado, err := service.GetOneWithContext(ctx, 2)

	assert.Nil(t, err)
	assert.Equal(t, usuarioNuevo.Nombre, usuarioCargado.Nombre)
}

func TestGetOneServiceSQLMock(t *testing.T) {

	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "nombre", "apellido", "email", "edad", "altura", "activo", "fecha_creacion"})
	rows.AddRow(1, "pedro", "sanchez", "psanchez@hotmail.com", 43, 0, false, "")
	mock.ExpectQuery("SELECT id, nombre,apellido, edad FROM usuarios WHERE id = ?").WithArgs(1).WillReturnRows(rows)

	repo := NewRepositorySQLMock(db)

	service := NewServiceSQL(repo)

	usuarioCargado := service.GetOne(1)

	assert.Nil(t, err)
	assert.Equal(t, "pedro", usuarioCargado.Nombre)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestStoreServiceSQLTxdb(t *testing.T) {
	newUser := models.Usuario{
		Nombre:        "jose",
		Apellido:      "lopez",
		Email:         "jlopez@hotmail.com",
		Edad:          41,
		Altura:        1.68,
		Activo:        true,
		FechaCreacion: "8/05/1966",
	}

	db, err := db.InitDb()
	assert.Nil(t, err)

	repo := NewRepositorySQLMock(db)
	defer db.Close()

	service := NewServiceSQL(repo)

	usuarioCreado, _ := service.Store(newUser.Nombre, newUser.Apellido, newUser.Email, newUser.Edad, newUser.Altura, newUser.Activo, newUser.FechaCreacion)

	assert.Equal(t, newUser.Nombre, usuarioCreado.Nombre)
	fmt.Println(usuarioCreado)
}
