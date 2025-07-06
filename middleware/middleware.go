package middleware

import "fmt"

func sumar(a, b int) int {
	return a + b
}

func restar(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func dividir(a, b int) int {
	return a / b
}

func MiMiddleware() {
	resultado := operacionesMidd(sumar)(2, 3)
	fmt.Println("Resultado suma: ", resultado)

	resultado = operacionesMidd(restar)(7, 3)
	fmt.Println("Resultado de resta: ", resultado)

	resultado = operacionesMidd(multiplicar)(2, 3)
	fmt.Println("Resultado de multiplicación: ", resultado)

	resultado = operacionesMidd(dividir)(2, 3)
	fmt.Println("Resultado de división: ", resultado)

}

func operacionesMidd(f func(int, int) int) func(int, int) int {
	return func(a, b int) int {
		fmt.Println("Inicio de operación")
		return f(a, b)
	}
}
