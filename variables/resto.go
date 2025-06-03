package variables

// Declaración de paquetes
import (
	"fmt"
	"strconv"
	"time" // Paquete para funciones de hora
)

var Nombre string
var Estado bool
var Sueldo float32
var Empresa string
var Empleados int
var Fecha time.Time

func RestoVariables() {
	Nombre = "Jesús Vaca"
	Estado = true
	Sueldo = 1577.66
	Fecha = time.Now()
	Empresa = "Conalep"
	Empleados = 1000

	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
	fmt.Println(Empresa)
	fmt.Println(Empleados)
}

func ConviertoaTexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
