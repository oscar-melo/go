package main

import "fmt"

var Calculo func(int, int) int = func(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Printf("sumo 5 + 7 = %d \n", Calculo(5, 7))

	Calculo = func(num1 int, num2 int) int {

		return num1 - num2
	}

	fmt.Printf("restamos 7 + 5 = %d \n", Calculo(7, 5))

	tablaDel := 2

	miTabla := Tabla(tablaDel)

	for i := 0; i < 10; i++ {
		fmt.Println(miTabla())
	}
}

func Tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {

		secuencia += 1
		return numero * secuencia

	}

}
