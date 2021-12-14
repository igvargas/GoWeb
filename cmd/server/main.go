package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/igvargas/GoWeb/cmd/server/handler"
	usuario "github.com/igvargas/GoWeb/internal/usuarios"
	store "github.com/igvargas/GoWeb/pkg/store"
	"github.com/joho/godotenv"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"github.com/swaggo/swag/example/basic/docs"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones

// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("No se pudo abrir el archivo .env")
	}

	router := gin.Default()

	db := store.New(store.FileType, "./usuarioSalid.json")

	repo := usuario.NewRepository(db)
	service := usuario.NewService(repo)
	controller := handler.NewUsuario(service)

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("/usuarios/get", controller.GetAll())
	router.POST("/usuarios/add", controller.Store())
	router.PUT("/usuarios/:id", controller.Update())
	router.PATCH("/usuarios/:id", controller.UpdateNombre())
	router.DELETE("/usuarios/:id", controller.Delete())

	router.Run(":8000")
}
