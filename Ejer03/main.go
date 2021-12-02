package main

import "fmt"

var estado bool
var numero int = 10

func main() {
	estado = false
	if estado = true; estado {
		fmt.Println("Hola Mundo")
	}

	switch numero {
	case 4:
		fmt.Println("Numero 4")
		break
	case 5:
		fmt.Println("Numero 5")
		break
	default:
		fmt.Println("Otro valor")
		break
	}

}
