package arreglosslices

import "fmt"

var tabla [10]int = [10]int{10, 0, 59}
var matriz [4][3]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[3] = 54
	tabla[6] = 7

	tabla2 := [10]string{"Pablo", "Juan"}
	fmt.Println(tabla)
	fmt.Println(tabla2)

	for i := 0; i < len(tabla); i++ {
		fmt.Println(tabla[i])
	}
	matriz[3][2] = 15
	fmt.Println(matriz)
}
