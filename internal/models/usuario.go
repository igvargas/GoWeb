package models

type Usuario struct {
	ID            int     `json:"id"`
	Nombre        string  `json:"nombre"`
	Apellido      string  `json:"apellido"`
	Email         string  `json:"email"`
	Edad          int     `json:"edad"`
	Altura        float64 `json:"altura"`
	Activo        bool    `json:"activo"`
	FechaCreacion string  `json:"fecha_creacion"`
	Domicilio     Ciudad  `json:"domicilio"`
}

type Ciudad struct {
	ID           int    `json:"id"`
	NombreCiudad string `json:"nombre_ciudad"`
	NombrePais   string `json:"nombre_pais"`
}
