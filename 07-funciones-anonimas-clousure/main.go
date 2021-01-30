package main

import "fmt"

// funcion anonima en una variable
var calculo func(int, int) int = func(num1, num2 int) int {
	return num1 + num2
}

func main() {
	// modificando el retorno de la funcion anonima. similar a la sobrecarga de funciones
	fmt.Println(calculo(5, 7))
	calculo = func(i1, i2 int) int {
		return i1 - i2
	}
	fmt.Println(calculo(6, 4))

	calculo = func(i1, i2 int) int {
		return i1 / i2
	}
	fmt.Println(calculo(90, 5))

	operaciones()

	// closures
	tablaDel := 2
	miTabla := tabla(tablaDel)
	for i := 1; i < 11; i++ {
		fmt.Println(miTabla())
	}

}

// funcion anonima declarada de otra forma
func operaciones() {
	resultado := func() int {
		var a int = 23
		var b int = 13
		return a + b
	}

	fmt.Println(resultado())
}

// closure isolacion de codigo.
func tabla(valor int) func() int {
	num := valor
	secuencia := 0
	return func() int {
		secuencia++
		return num * secuencia
	}
}
