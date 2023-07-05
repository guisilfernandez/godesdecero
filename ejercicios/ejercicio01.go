package ejercicios

import (
	"strconv"
)

func ResultadosEjercicio01(valor string) (int, string) {
	numero, err := strconv.Atoi(valor)
	if err != nil {
		return 0, "Hubo un error " + err.Error()
	}
	if numero > 100 {
		return numero, "Es mayor que 100"
	} else {
		return numero, "Es menor que 100"
	}
}
