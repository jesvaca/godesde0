package main

/*
	Este es mi primer programa en Golang
	Desarrollado el 1 de Junio del 2025
	Excelente lenguaje
*/
import (
	"fmt"
	"runtime"

	"github.com/jesvaca/godesde0/ejercicios"
	"github.com/jesvaca/godesde0/variables"
)

func main() {
	variables.MuestroEnteros()
	variables.RestoVariables()
	estado, texto := variables.ConviertoaTexto(1588)
	fmt.Println(estado)
	fmt.Println(texto)

	//Obtenemos información del sistema operativo gracias al paquete runtime (declarado en import)
	os := runtime.GOOS
	fmt.Println("Sistema operativo (fuera del if): ", os) // imprimir la versión del sistema operativo

	// Uso del if-else
	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows, es ", os)
	} else {
		fmt.Println("Esto es Windows")
	}

	// Uso del switch
	fmt.Println("Mediante el SWITCH")
	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "darwin":
		fmt.Println("Esto es Darwin")
	case "windows":
		fmt.Printf("Esto es Windows %s \n", os)
	case "Linux.":
		fmt.Println("Esto es Linux.")
	case "OS X.":
		fmt.Println("Esto es OS X.")
	default:
		fmt.Printf("%s \n", os)
	}

	cadenaNumero := "500"
	numero, mensaje := ejercicios.ConvierteStringANumero(cadenaNumero)
	fmt.Println(numero)
	fmt.Println(mensaje)
}
