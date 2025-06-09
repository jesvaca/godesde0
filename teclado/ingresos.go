package teclado

import (
	"bufio" // Paquete de entrada de datos por varios dispositivos
	"fmt"
	"os"
	"strconv"
)

var numero1 int
var numero2 int
var leyenda string
var err error

func IngresoNumeros() {
	/* 	PRIMERA OPCIÓN DE CAPTURA DE TECLADO usando SCANLN()

	fmt.Println("Ingrese numero 1: ")
	fmt.Scanln(&numero1)
	fmt.Println("Ingrese numero 2: ")
	fmt.Scanln(&numero2)
	fmt.Println("Ingrese leyenda:")
	fmt.Scanln(&leyenda)
	fmt.Println(leyenda, numero1*numero2)

	fmt.Println("NUEVAS CAPTURAS:")
	*/
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingresa de nuevo el número 1: ")
	if scanner.Scan() {
		numero1, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El número 1 ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingresa de nuevo el número 2: ")
	if scanner.Scan() {
		numero2, err = strconv.Atoi(scanner.Text())
		if err != nil {
			panic("El número 2 ingresado es incorrecto " + err.Error())
		}
	}

	fmt.Println("Ingresa de nuevo la leyenda: ")
	if scanner.Scan() {
		leyenda = scanner.Text()
	}
	fmt.Println(leyenda, numero1*numero2)

}
