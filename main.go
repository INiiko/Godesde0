package main

import (
	"fmt"

	"github.com/Godesde0/ejercicios"
)

// var Numero = 500 ConvNumerico("500")
func main() {
	// estado, texto := variables.ConviertoaTexto(1588)
	// fmt.Println(estado)
	// fmt.Println(texto)

	numero, texto := ejercicios.ConvNumerico("50")
	fmt.Println(numero)
	fmt.Println(texto)

}
