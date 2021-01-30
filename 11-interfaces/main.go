package main

import "fmt"

type SerVivo interface {
	estaVivo() bool
}

type humano interface {
	respirar()
	pensar()
	comer()
	genero() string
	estaVivo() bool
}

type animal interface {
	respirar()
	comer()
	EsCarnivoro() bool
	estaVivo() bool
}

type vegetal interface {
	ClasificacionVegeta() string
}

// genero humano
type hombre struct {
	edad       int
	altura     float32
	peso       float32
	respirando bool
	pensando   bool
	comiendo   bool
	sexo       bool
	vivo       bool
}

// genero humano
type mujer struct {
	hombre
}

func (h *hombre) respirar() {
	h.respirando = true
}
func (h *hombre) comer() {
	h.comiendo = true
}
func (h *hombre) pensar() {
	h.pensando = true
}
func (h *hombre) genero() string {
	if h.sexo {
		return "masculino"
	}
	return "femenino"
}

func (h *hombre) estaVivo() bool {
	return h.vivo
}

func HumanosRespirando(hu humano) {
	hu.respirar()
	fmt.Printf("soy %s, y estoy respirando \n", hu.genero())
}

// genero animal

type perro struct {
	respirando bool
	comiendo   bool
	carnivoro  bool
	vivo       bool
}

func (p *perro) respirar() {
	p.respirando = true
}
func (p *perro) comer() {
	p.comiendo = true
}
func (p *perro) EsCarnivoro() bool {
	return p.carnivoro
}

func (p *perro) estaVivo() bool {
	return p.vivo
}

func AnimalesRespirar(an animal) {
	an.respirar()
	fmt.Println("soy un animal y estoy repirando")
}

func AnimalesCarnivoros(an animal) int {
	if an.EsCarnivoro() {
		return 1
	}
	return 0
}

// SerVivo
func estoyVivo(v SerVivo) bool {
	return v.estaVivo()
}

func main() {
	Eduardo := new((hombre))
	Eduardo.sexo = true
	HumanosRespirando(Eduardo)
	Maria := new((mujer))
	HumanosRespirando(Maria)

	totalCarnivoros := 0
	dogo := new(perro)
	dogo.carnivoro = true
	dogo.vivo = true
	AnimalesRespirar(dogo)
	totalCarnivoros += AnimalesCarnivoros((dogo))
	fmt.Println(totalCarnivoros)
	fmt.Printf("Estoy vivo = %t\n", estoyVivo(dogo))
}
