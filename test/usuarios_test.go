package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/igvargas/GoWeb/cmd/server/handler"
	usuario "github.com/igvargas/GoWeb/internal/usuarios"
	"github.com/igvargas/GoWeb/pkg/store"
	"github.com/igvargas/GoWeb/pkg/web"
	"github.com/stretchr/testify/assert"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")

	// We want to make sure the token is set, bail if not
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}

	return func(c *gin.Context) {
		token := c.GetHeader("token")

		if token == "" {
			respondWithError(c, 401, "API token required")
			return
		}

		if token != requiredToken {
			respondWithError(c, 401, "Invalid API token")
			return
		}

		c.Next()
	}
}

func createServer() *gin.Engine {
	router := gin.Default()
	_ = os.Setenv("TOKEN", "123456")
	db := store.New(store.FileType, "./usuariosSalidaTest.json")
	repo := usuario.NewRepository(db)
	service := usuario.NewService(repo)
	controller := handler.NewUsuario(service)

	router.Use(TokenAuthMiddleware())

	router.GET("/usuarios/get", controller.GetAll())
	router.POST("/usuarios/add", controller.Store())
	router.PUT("/usuarios/:id", controller.Update())
	router.PATCH("/usuarios/:id", controller.UpdateNombre())
	router.DELETE("/usuarios/:id", controller.Delete())

	return router
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

func Test_GetUsuarios(t *testing.T) {
	router := createServer()

	req, rr := createRequestTest(http.MethodGet, "/usuarios/get", "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta web.Response

	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	assert.Equal(t, 200, respuesta.Code)
	assert.Nil(t, err)
}

func Test_StoreUsuarios(t *testing.T) {
	router := createServer()

	nuevoUsuario := usuario.Usuario{
		Nombre:        "cristian",
		Apellido:      "lopez",
		Email:         "clopez@hotmail.com",
		Edad:          41,
		Altura:        1.78,
		Activo:        true,
		FechaCreacion: "17/05/2001",
	}

	fmt.Println(nuevoUsuario)
	dataNueva, _ := json.Marshal(nuevoUsuario)
	fmt.Println(string(dataNueva))
	req, rr := createRequestTest(http.MethodPost, "/usuarios/add", string(dataNueva))

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

	var respuesta usuario.Usuario
	//fmt.Printf(string(rr.Body.Bytes()))
	// fmt.Printf("el nombre del user es: %v\n", string(respuesta))
	err := json.Unmarshal(rr.Body.Bytes(), &respuesta)
	//fmt.Printf("el nombre del user es: %v\n", respuesta)
	assert.Equal(t, "cristian", respuesta.Nombre)
	assert.Nil(t, err)

	////////////////////////////////
	delete := fmt.Sprintf("/usuarios/%d", respuesta.ID)
	req, rr = createRequestTest(http.MethodDelete, delete, "")

	router.ServeHTTP(rr, req)

	assert.Equal(t, 200, rr.Code)

}
