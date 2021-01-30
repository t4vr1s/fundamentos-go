package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	leerArchivo()
	leerOtroArchivo()
	grabarArchivo()
	graboOtroArchivo()
}

func leerArchivo() {
	archivo, err := ioutil.ReadFile("12-manejo-de-archivos/archivo.txt")
	if err != nil {
		fmt.Println("hubo un error", err)
		// return nil, err
	}
	fmt.Println(string(archivo))
}

func leerOtroArchivo() {
	archivo, err := os.Open("12-manejo-de-archivos/archivo.txt")
	if err != nil {
		fmt.Println("hubo un error", err)
	}
	escaner := bufio.NewScanner(archivo)
	for escaner.Scan() {
		registro := escaner.Text()
		fmt.Printf("linea > " + registro + "\n")
	}
	archivo.Close()
}

func grabarArchivo() {
	archivo, err := os.Create("12-manejo-de-archivos/nuevoArchivo.txt")
	if err != nil {
		fmt.Println("hubo un error")
		archivo.Close()
		return
	}
	fmt.Fprintln(archivo, "esta es una linea nueva")
	archivo.Close()
}

func graboOtroArchivo() {
	fileName := "12-manejo-de-archivos/nuevoArchivo.txt"
	if Append(fileName, "\nesta es la segunda linea") == false {
		fmt.Println("error en la segunda linea")
	}
}

func Append(file string, text string) bool {
	archivo, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("hubo un error", err)
		archivo.Close()
		return false
	}

	_, err = archivo.WriteString(text)
	if err != nil {
		fmt.Println("hubo un error", err)
		archivo.Close()
		return false
	}

	archivo.Close()
	return true
}
