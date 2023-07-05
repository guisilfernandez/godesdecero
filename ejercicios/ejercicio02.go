package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var numero int
var err error

func PideNumero() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese el número :")
	if scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El dato ingresado es incorrecto " + err.Error())
		}
		for i := 1; i < 11; i++ {
			fmt.Printf("%d * %d = %d\n", numero, i, numero*i)
		}
	}

}
