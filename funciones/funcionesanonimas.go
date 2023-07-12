package funciones

import "fmt"

var numero3 int = 32
var numero4 int = 243

func Calculos() {

	suma := func(numero1 int, numero2 int) int {
		return numero1 + numero2 + numero3 + numero4
	}

	fmt.Println(suma(5, 10))

	suma = func(numero1 int, numero2 int) int {
		return numero1 * numero2 * numero3
	}
	fmt.Println(suma(5, 10))
}
