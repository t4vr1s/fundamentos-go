package main

import (
	"fmt"
	"time"

	us "./user"
)

type pepe struct {
	us.Usuario
}

func main() {
	user := new(pepe)
	user.AltaUsuario(1, "eduardo escobar", time.Now(), true)
	fmt.Println(user.Usuario)
}
