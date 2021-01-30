package main

import "fmt"

// declarando variables de ambito global
var numero int
var texto string
var estado bool

func main() {
	// declarando variables con asignacion corta
	numero2 := 3
	var moneda int16 = 234
	// asignando un valor a la variable
	numero2 = int(moneda)
	fmt.Println(numero2)
	// conversión

	// variables encadenadas
	var num1, num2, num3 int
	num4, num5, num6, texto, verdad := 2, 4, 6, "eduardo", true
	fmt.Println(num4, num5, num6, texto, verdad)
	// valor zero
	fmt.Println(num1, num2, num3)
}
