package main

import "fmt"

var estado bool

func main() {
	// al if se le puede iniciar el valor en la misma linea de la condición
	estado = true
	if estado = false; estado {
		fmt.Println("es verdadero")
	} else {
		fmt.Println("es falso")
	}

	// al switch se le puede iniciar el valor en la misma linea de la expresión y no tiene break
	switch numero := 10; numero {
	case 1:
		fmt.Println(numero)
	case 2:
		fmt.Println(numero)
	case 3:
		fmt.Println(numero)
	default:
		fmt.Println("no es 1, 2 o 3 ")
	}
}
