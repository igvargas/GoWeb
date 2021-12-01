package main

import (
	"github.com/gin-gonic/gin"
	"github.com/igvargas/GoWeb/cmd/server/handler"
	usuario "github.com/igvargas/GoWeb/internal/usuarios"
)

func main() {
	router := gin.Default()

	repo := usuario.NewRepository()
	service := usuario.NewService(repo)
	controller := handler.NewUsuario(service)

	router.GET("/usuarios/get", controller.GetAll())
	router.POST("/usuarios/add", controller.Store())
	router.PUT("/usuarios/:id", controller.Update())
	router.PATCH("/usuarios/:id", controller.UpdateNombre())
	router.DELETE("/usuarios/:id", controller.Delete())

	router.Run(":8000")
}
