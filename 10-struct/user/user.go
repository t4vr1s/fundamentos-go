package user

import "time"

type Usuario struct {
	Id        int
	Nombre    string
	FechaAlta time.Time
	Estado    bool
}

func (this *Usuario) AltaUsuario(id int, nombre string, fechaalta time.Time, estado bool) {
	this.Id = id
	this.Nombre = nombre
	this.FechaAlta = fechaalta
	this.Estado = estado
}
