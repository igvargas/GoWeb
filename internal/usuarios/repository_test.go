package internal

import (
	"encoding/json"
	"testing"

	store "github.com/igvargas/GoWeb/pkg/store"
	"github.com/stretchr/testify/assert"
)

var usr string = `[
	{"id": 1,"nombre": "josefa","apellido": "Maylor","email": "lmaylor0@dagondesign.com","edad": 30,"altura": 177.1,"activo": false,"fecha_creacion": "8/20/2021"},
    {"id": 2,"nombre": "Lenard","apellido": "Maylor","email": "lmaylor0@dagondesign.com","edad": 30,"altura": 177.1,"activo": false,"fecha_creacion": "8/20/2021"}]`

var sliceDeUsuarios []Usuario = []Usuario{{
	ID:            1,
	Nombre:        "cristian",
	Apellido:      "lopez",
	Email:         "clopez@hotmail.com",
	Edad:          41,
	Altura:        1.78,
	Activo:        true,
	FechaCreacion: "17/05/2001",
},
	{
		ID:            1,
		Nombre:        "cristian",
		Apellido:      "lopez",
		Email:         "clopez@hotmail.com",
		Edad:          41,
		Altura:        1.78,
		Activo:        true,
		FechaCreacion: "17/05/2001",
	},
}

type StubStore struct {
}

type MockStore struct {
	useUpdateNombre bool
}

func (s *StubStore) Read(data interface{}) error {
	sliceDeBytes, err := json.Marshal(sliceDeUsuarios)
	if err != nil {
		return err
	}
	return json.Unmarshal(sliceDeBytes, &data)
}

func (s *StubStore) Write(data interface{}) error {
	return nil
}

func (m *MockStore) Read(data interface{}) error {
	m.useUpdateNombre = true
	return json.Unmarshal([]byte(usr), &data)

}

func (m *MockStore) Write(data interface{}) error {
	return nil
}

func TestGetAll(t *testing.T) {
	// Arrange
	store := StubStore{}
	repo := NewRepository(&store)

	// Act
	users, err := repo.GetAll()
	//var expected []Usuario
	//json.Unmarshal([]byte(usr), &expected)

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, 2, len(users))
}

func TestUpdateName(t *testing.T) {
	// Arrange
	store := MockStore{}
	repo := NewRepository(&store)

	nombreExpected := "Pepe"

	// Act
	userUpdated, _ := repo.UpdateNombre(1, "Pepe")

	// Assert
	assert.Equal(t, nombreExpected, userUpdated.Nombre)
	assert.True(t, store.useUpdateNombre)
}

func TestUpdateNameError(t *testing.T) {
	// Arrange
	store := MockStore{}
	repo := NewRepository(&store)

	// Act
	_, err := repo.UpdateNombre(4, "Pepe")

	// Assert
	assert.Error(t, err)
}

// func TestUpdateNameRead(t *testing.T) {
// 	// Arrange
// 	store := MockStore{}

// 	// Act

// 	// Assert
// 	assert.True(t, store.useUpdateNombre)
// }

func TestGetAllRepositoryMock(t *testing.T) {
	// Arrange
	dataByte := []byte(usr)
	var usrExpected []Usuario
	json.Unmarshal(dataByte, &usrExpected)

	dbMock := store.Mock{Data: dataByte}
	storeStub := store.FileStore{Mock: &dbMock}
	repo := NewRepository(&storeStub)

	// Act
	misUser, _ := repo.GetAll()

	// Assert
	assert.Equal(t, usrExpected, misUser)
}
