package ejercicios

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func TablaMultiplicar() string {
	var numero int
	var err error
	var texto string
	scanner := bufio.NewScanner(os.Stdin) // Iniciamos la variable de captura de número
	fmt.Print("Ingresa el número a multiplicar: ")
	for scanner.Scan() {
		numero, err = strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Dato incorrecto...")
			fmt.Print("Ingresa el número a multiplicar: ")
		} else {
			break
		}
	}
	for i := 1; i <= 10; i++ {
		// Dos opciones para el depliegue de la tabla
		// resultado = numero * i
		// fmt.Println(numero, "x", i, "=", resultado)
		// fmt.Printf("%d x %d = %d\n", numero, i, numero*i)
		texto += fmt.Sprintf("%d x %d = %d\n", numero, i, numero*i)
	}
	return texto
}
