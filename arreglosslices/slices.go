/* SLICES arreglos dinámicos */
package arreglosslices

import "fmt"

var tablaSlice []int = []int{2, 5, 4}
var matrizSlice [][]int = [][]int{{1, 2, 3}, {4, 5, 6}}

func MuestroSlices() {
	fmt.Println(tablaSlice)
	tablaSlice = append(tablaSlice, 8)
	fmt.Println(tablaSlice)

}

func Capacidad() {
	elementos := make([]int, 5, 20)
	fmt.Printf("Largo %d, Capacidad %d\n", len(elementos), cap(elementos))
	fmt.Println(elementos)
	nums := make([]int, 0, 0)
	for i := 0; i < 100; i++ {
		nums = append(nums, i)
	}
	fmt.Printf("\nLargo %d, Capacidad %d", len(nums), cap(nums))
}
