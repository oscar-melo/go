package main

import (
	"bufio"
	"fmt"
	"os"
)

var numero int
var numero2 int
var leyenda string
var resultado int

func main() {
	fmt.Println("Ingrese el numero 1 ")
	fmt.Scanf("%d", &numero)

	fmt.Println("Ingrese el numero 2 ")
	fmt.Scanf("%d", &numero2)

	fmt.Println("Descripcion: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		leyenda = scanner.Text()

	}
	resultado = numero + numero2
	fmt.Println(leyenda, resultado)

}
