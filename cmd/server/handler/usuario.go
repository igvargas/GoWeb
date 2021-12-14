package handler

import (
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
	usuario "github.com/igvargas/GoWeb/internal/usuarios"
	"github.com/igvargas/GoWeb/pkg/web"
)

type request struct {
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int     `json:"edad"`
	Altura        float64 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fecha_creacion"`
}

type Usuario struct {
	service usuario.Service
}

func NewUsuario(ser usuario.Service) *Usuario {
	return &Usuario{service: ser}
}

// ListUsers godoc
// @Summary List usuarios
// @Tags Usuario
// @Description get usuarios
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Router /usuarios/get [get]
func (usr *Usuario) GetAll() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		/////////
		ok := ValidarToken(ctx)
		if !ok {
			return
		}
		/////////
		usuarios, err := usr.service.GetAll()

		if err != nil {
			ctx.JSON(200, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error %v", err)))
			//ctx.String(400, "Hubo un error %v", err)
		} else {
			ctx.JSON(200, web.NewResponse(200, usuarios, ""))
		}
	}
}

// StoreUsers godoc
// @Summary Store usuario
// @Tags Usuario
// @Description store usuario
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param persona body request true "Usuario to store"
// @Success 200 {object} web.Response
// @Router /usuarios/add [post]
func (controller *Usuario) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ok := ValidarToken(ctx)
		if !ok {
			return
		}
		var usu request

		err := ctx.ShouldBindJSON(&usu)

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Hubo un error al querer cargar una persona %v", err)))
		} else {
			response, err := controller.service.Store(usu.Nombre, usu.Apellido, usu.Email, usu.Edad, usu.Altura, usu.Activo, usu.FechaCreacion)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("No se pudo cargar la persona %v", err)))
			} else {
				ctx.JSON(200, web.NewResponse(200, response, ""))
			}
		}

	}
}

// UpdateUser godoc
// @Summary Update usuario
// @Tags Usuario
// @Description update usuario
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param persona body request true "Usuario update to store"
// @Success 200 {object} web.Response
// @Router /usuarios/:id [put]
func (controller *Usuario) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		ok := ValidarToken(ctx)
		if !ok {
			return
		}

		var usr request
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("El id es invalido")))
		}

		err = ctx.ShouldBindJSON(&usr)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Error en el body")))
		} else {
			usuarioActualizado, err := controller.service.Update(int(id), usr.Nombre, usr.Apellido, usr.Email, usr.Edad, usr.Altura, usr.Activo, usr.FechaCreacion)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			} else {
				ctx.JSON(200, web.NewResponse(200, usuarioActualizado, ""))
			}
		}

	}
}

// UpdateUserName godoc
// @Summary Update nombre usuario
// @Tags Usuario
// @Description update nombre usuario
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param persona body request true "Usuario update nombre to store"
// @Success 200 {object} web.Response
// @Router /personas/:id [patch]
func (controller *Usuario) UpdateNombre() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ok := ValidarToken(ctx)
		if !ok {
			return
		}
		var usr request

		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("El id es invalido")))
		}
		err = ctx.ShouldBindJSON(&usr)
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("Error en el body")))
		} else {
			if usr.Nombre == "" {
				ctx.JSON(404, web.NewResponse(400, nil, fmt.Sprintf("El nombre no puede estar vacío")))
				return
			}
			usuarioActualizado, err := controller.service.UpdateNombre(int(id), usr.Nombre)
			if err != nil {
				ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
			} else {
				ctx.JSON(200, web.NewResponse(200, usuarioActualizado, ""))
			}
		}
	}
}

// DeleteUser godoc
// @Summary Delete usuario
// @Tags Usuario
// @Description delete usuario
// @Accept  json
// @Produce  json
// @Param token header string true "token"
// @Param persona body request true "Delete Usuario"
// @Success 200 {object} web.Response
// @Router /usuarios/:id [Delete]
func (controller *Usuario) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ok := ValidarToken(ctx)
		if !ok {
			return
		}
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, fmt.Sprintf("El id es invalido")))
		}
		err = controller.service.Delete(int(id))
		if err != nil {
			ctx.JSON(400, web.NewResponse(400, nil, err.Error()))
		} else {
			ctx.JSON(200, web.NewResponse(200, nil, fmt.Sprintf("La persona %d ha sido eliminada", id)))
		}
	}
}

func ValidarToken(ctx *gin.Context) bool {
	token := ctx.GetHeader("token")
	tokenENV := os.Getenv("TOKEN")
	if token == "" {
		ctx.JSON(200, web.NewResponse(400, nil, "Falta token"))
		//ctx.String(400, "Falta token")
		return false
	}

	if token != tokenENV {
		ctx.JSON(200, web.NewResponse(404, nil, "Token incorrecto"))
		//ctx.String(404, "Token incorrecto")
		return false
	}
	return true
}
