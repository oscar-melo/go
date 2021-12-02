package main

import (
	"fmt"
	"time"

	us "ejer07/user"
)

var tabla = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}

func main() {
	for _, numero := range tabla {
		fmt.Println(numero)
	}

	fmt.Println(tabla[3:4])

	fmt.Println(len(tabla))
	fmt.Println(cap(tabla))

	paises := make(map[string]string)
	paises["1"] = "uno"
	paises["2"] = "dos"

	fmt.Println(paises)

	user := new(us.Usuario)

	user.Id = 10
	user.Nombre = "Oscar"
	user.FechaAlta = time.Now()
	user.Status = true

	fmt.Println(user)

	pepe := new(Pepe)
	pepe.Nombre = "PEpe"

	fmt.Println(pepe)

}

type Pepe struct {
	us.Usuario
}
