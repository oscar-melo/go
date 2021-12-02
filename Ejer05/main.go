package main

import "fmt"

func main() {

	numero, estado := uno(5)
	fmt.Println(numero)
	fmt.Println(estado)

	fmt.Println(dos(6))

	fmt.Println(calculo(5, 4, 6, 3, 2))

}

func uno(numero int) (int, bool) {
	return numero * 2, true
}

func dos(numero int) (z int) {
	z = numero * 2
	return
}

func calculo(numero ...int) int {

	total := 0
	for _, num := range numero {
		total += num
	}

	return total

}
