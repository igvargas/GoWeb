package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/igvargas/GoWeb/cmd/server/handler"
	usuario "github.com/igvargas/GoWeb/internal/usuarios"
	store "github.com/igvargas/GoWeb/pkg/store"
	"github.com/joho/godotenv"
)

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

	router.GET("/usuarios/get", controller.GetAll())
	router.POST("/usuarios/add", controller.Store())
	router.PUT("/usuarios/:id", controller.Update())
	router.PATCH("/usuarios/:id", controller.UpdateNombre())
	router.DELETE("/usuarios/:id", controller.Delete())

	router.Run(":8000")
}
