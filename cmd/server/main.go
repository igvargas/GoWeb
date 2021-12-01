package main

import (
	"github.com/gin-gonic/gin"
	u "github.com/igvargas/GoWeb/cmd/server/handler"
	usuario "github.com/igvargas/GoWeb/internal/usuarios"
)

func main() {
	router := gin.Default()

	repo := usuario.NewRepository()
	service := usuario.NewService(repo)
	controller := u.NewUsuario(service)

	router.GET("/usuarios/get", controller.GetAll())
	router.POST("/usuarios/add", controller.Store())

	router.Run(":8000")
}
