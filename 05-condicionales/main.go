package main

import "fmt"

func main() {
	i := 0
	// for tipo while
	for i < 10 {
		fmt.Println("desde el for tipo while", i)
		i++
	}
	// for tradicional
	for i := 0; i < 10; i++ {
		fmt.Println("desde el for tradicional", i)
	}

	// for infinito ¿do while? con break
	numero := 0
	for {
		fmt.Println("for break")
		fmt.Println("ingrese numero secreto")
		fmt.Scanf("%d", &numero)
		if numero == 100 {
			fmt.Println("correcto, adios")
			break
		}
	}

	// for con continue
	numero2 := 0
	for numero2 < 10 {
		fmt.Println("el valor de i es: ", numero2)
		if numero2 == 5 {
			fmt.Printf("multiplicamos por 2 \n")
			numero2 *= 2
			continue
		}
		numero2++
	}

	var num int

RUTINA:
	fmt.Println("salió del for")
	for num < 10 {
		if num == 4 {
			num += 2
			fmt.Println("voy a rutina sumando 2 a num")
			goto RUTINA
		}
		fmt.Printf("el valor de num es %d\n", num)
		num++
	}
}
