package main

import (
	"github.com/jesvaca/godesde0/mapas"
)

/*
	Este es mi primer programa en Golang
	Desarrollado el 1 de Junio del 2025
	Excelente lenguaje
*/

func main() {
	//var tablaMul string
	/*
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

		cadenaNumero := "100"
		numero, mensaje := ejercicios.ConvierteStringANumero(cadenaNumero)
		fmt.Println("El número es: ", numero)
		fmt.Println(mensaje)

		teclado.IngresoNumeros() */

	// Tema ciclos for
	// ciclos.Ciclos(10, 30)
	/*
		tablaMul := ejercicios.TablaMultiplicar()
		fmt.Println("Tabla de multiplicar")
		fmt.Println("===================")
		fmt.Println(tablaMul)
		// files.GrabaTabla(tablaMul)
		files.SumaTabla(tablaMul)
	*/
	// files.LeoArchivo()
	// funciones.Calculos()
	// funciones.LlamarClosure()
	// funciones.Exponencia(3)
	// arreglosslices.MuestroArreglos()
	// arreglosslices.MuestroSlices()
	// arreglosslices.Capacidad()
	mapas.MostrarMapas()
}
