package main

import (
	"fmt"
	"github-cursos/godesdecero/ejercicios"
)

func main() {
	//estado, texto := variables.ConviertoaTexto(1588)
	//fmt.Println(estado)
	//fmt.Println(texto)

	/*if os := runtime.GOOS; os == "Linux." || os == "OS X." || os == "darwin" {
		fmt.Println("Esto no es Windows, esto es", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s \n", os)
	}*/

	valor, texto := ejercicios.ResultadosEjercicio01("1985")
	fmt.Println("Enviado", valor)
	fmt.Println(texto)
}
