package main

import "fmt"

var result int

func main() {
	fmt.Println("inicio")
	result = operacionesMidd(sumar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(restar)(2, 3)
	fmt.Println(result)
	result = operacionesMidd(multiplicar)(2, 3)
	fmt.Println(result)
}

func sumar(a, b int) int {
	return a + b
}
func restar(a, b int) int {
	return a - b
}
func multiplicar(a, b int) int {
	return a * b
}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(i1, i2 int) int {
		fmt.Println("inicio de operación")
		return f(i1, i1)
	}

}
