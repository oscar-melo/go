package main

import "fmt"

var numero int = 1
var texto string
var status bool

func main() {
	numero4, numero, texto := 3, 4, "oscar"

	var moneda float32 = 4.2342

	numero = int(moneda)
	fmt.Println(numero4)
	fmt.Println(numero)
	fmt.Println(texto)

	fmt.Println("Cambio")
	fmt.Println(numero)
}
