package main

import "fmt"

// declarando un arreglo
var tabla [10]int

// declarando una matriz
var matriz [2][2]int

func main() {
	// asignar valores
	tabla[0] = 1
	tabla[5] = 15
	fmt.Println(tabla)
	array := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	fmt.Printf("%T\n", array)
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}

	//asignando valor a una matriz
	matriz[1][1] = 1
	fmt.Println(matriz)

	// crear un slice por asignación
	mySlice := []int{2, 4, 6}
	fmt.Println(mySlice)
	variante2()
	variante3()
	variante4()
}

func variante2() {
	// creando un arreglo
	elementos := [5]int{1, 2, 4, 5, 6}
	// creando un slice por asignación desde un array desde el elemento 3 hasta el final
	porcion := elementos[3:]
	// porcion := elementos[:3]
	// porcion := elementos[1:3]
	fmt.Println(porcion)
}

func variante3() {
	// crear un slice con make
	elementos := make([]int, 5, 20)
	fmt.Printf("largo %d, capacidad %d\n", len(elementos), cap(elementos))
}

func variante4() {
	nums := make([]int, 0, 0)
	for i := 0; i < 130; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("largo %d, capacidad %d\n ", len(nums), cap(nums))
}
