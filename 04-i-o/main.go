package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero1 int
var numero2 int
var resultado int
var leyenda string

func main() {
	fmt.Println("ingrese numero 1")
	fmt.Scanf("%d", &numero1)
	// fmt.Scanln("%d", &numero1) en so windows usar 'Scanln'
	fmt.Println("ingrese numero 2")
	fmt.Scanf("%d", &numero2)
	// fmt.Printf("%d/n", numero1+numero2)
	fmt.Println("Descripcion: ")
	escaner := bufio.NewScanner(os.Stdin)
	if escaner.Scan() {
		leyenda = escaner.Text()
	}
	resultado = numero1 + numero2
	fmt.Println(leyenda, resultado)
}
