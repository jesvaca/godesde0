package funciones

import "fmt"

func Exponencia(valor int) {
	if valor == 0 {
		fmt.Println("El valor debe ser mayor de 0")
		return
	} else {
		if valor > 1000000000 {
			return
		}
		fmt.Println(valor)
		Exponencia(valor * 2)
	}
}
