package main

import "fmt"

func main() {
	// creando un mapa
	paises := make(map[string]string)
	fmt.Println(paises)
	// asignando un elemento
	paises["chile"] = "santiago"
	paises["argentina"] = "buenos aires"
	// obtener el elemento
	fmt.Println(paises["chile"])
	fmt.Println(paises)

	// creando un mapa otra forma
	campeonato := map[string]int{
		"colo-colo":  39,
		"catolica":   39,
		"u de chile": 39,
		"iberia":     44,
	}
	fmt.Println(campeonato)
	// agregando un elemento
	campeonato["osorno"] = 22
	fmt.Println(campeonato)
	// modificar un elemento
	campeonato["iberia"] = 100
	fmt.Println(campeonato)
	// eliminar un elemento
	delete(campeonato, "iberia")
	fmt.Println(campeonato)
	// recorrer un mapa
	for equipo, puntaje := range campeonato {
		fmt.Printf("el equipo %s, tiene: %d puntos \n", equipo, puntaje)
	}

	// saber si existe un elemento en el mapa
	puntaje, existe := campeonato[""]
	fmt.Printf("el puntaje del equipo es: %d, y existe: %v\n", puntaje, existe)
}
