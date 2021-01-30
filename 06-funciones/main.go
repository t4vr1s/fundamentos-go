package main

import "fmt"

func main() {
	fmt.Println(uno(5))
	num, estado := dos(2)
	fmt.Println(num, estado)

	fmt.Println(calculo(5, 46))
	fmt.Println(calculo(5, 46, 2, 78))
	fmt.Println(calculo(5))
	fmt.Println(calculo(5, 46, 17, 25, 24, 99, 78))
}

func uno(numero int) int {
	return numero * 2
}

// retorno de 2 valores
func dos(num int) (int, bool) {
	if num == 1 {
		return 5, false
	}
	return 10, true
}

// parametros variables
func calculo(numero ...int) int {
	total := 0
	for _, num := range numero {
		total += num
	}
	return total
}
