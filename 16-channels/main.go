package main

import (
	"fmt"
	"time"
)

func main() {
	canal1 := make(chan time.Duration)
	go bucle(canal1)
	fmt.Println("llegué has aquí")
	msg := <-canal1
	fmt.Println(msg)
}

func bucle(canal chan time.Duration) {
	inicio := time.Now()
	for i := 0; i < 100000000000; i++ {

	}

	final := time.Now()
	canal <- final.Sub(inicio)
}
