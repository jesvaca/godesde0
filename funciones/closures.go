package funciones

import "fmt"

func tabla(valor int) func() int {
	numero := valor
	secuencia := 0
	return func() int {
		secuencia++
		return numero * secuencia
	}
}

func LlamarClosure() {
	tabladel := 2
	closure := tabla(tabladel)
	for i := 1; i < 11; i++ {
		fmt.Println(tabladel, "x", i, "=", closure())
	}
}
