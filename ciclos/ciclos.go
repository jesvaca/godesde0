package ciclos

import (
	"fmt"
)

func Ciclos(inic, fin int) {
	// Formas FOR

	for i := inic; i <= fin; i += 5 {
		if i == 20 {
			continue // Al coincidir con el 20, se brinca todo el codigo
		}
		fmt.Print(i, " ")
	}

	/*
		i := 0
		for i < count {
			fmt.Println(i)
			i++
		}
	*/

}
