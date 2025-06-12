// FUNCIONES ANONIMAS
package funciones

import "fmt"

// Función anónima sobre la variable calculo
func Calculos() {

	// La función anónima se crea asignando sobre una variable, la cual se llama de otro punto de programa
	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}
	fmt.Println(calculo(2, 3))

	/* La función anónima que se definión sobre la variable calculo, se reescribe (SOBRECARGA)
	para procesar otros datos */
	calculo = func(numero1 int, numero2 int) int {
		return numero1 * numero2
	}
	fmt.Println(calculo(2, 3))

}
