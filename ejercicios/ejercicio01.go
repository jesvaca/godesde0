package ejercicios

import "strconv"

func ConvierteStringANumero(cadena string) (int, string) {
	//var numero int64
	// numero, err := strconv.ParseInt(cadena, 10, 64)     <- Una opción de cconversió
	numero, err := strconv.Atoi(cadena) // <- Otra opción de conversión
	if err != nil {
		return 0, "Hubo un error "
	}
	if numero > 100 {
		return numero, "Es mayor a 100"
	} else {
		return numero, "Es menor a 100"
	}
}
