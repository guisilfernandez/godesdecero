package main

import (
	"fmt"
	"github-cursos/godesdecero/variables"
)

func main() {
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)
}
