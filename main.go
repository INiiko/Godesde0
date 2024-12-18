package main

import (
	"fmt"
	"runtime"
	"\ejercicios"
)

func main() {
	/* estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	if os := runtime.GOOS; os == "linux" {
		fmt.Println("Esto no es Windows")
	} else {
		fmt.Println("Esto es Windows")
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es darwin")
	default:
		fmt.Printf("%s \n", os)
	}
*/
		numero, texto := ejercicios.ConvNumerico("500")
		fmt.println(numero)
		fmt.Println(texto)
}
