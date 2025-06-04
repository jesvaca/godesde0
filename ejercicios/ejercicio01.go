package ejercicios

import "strconv"

func ConvierteStringANumero(cadena string) (int64, string) {
	//var numero int64
	numero, err := strconv.ParseInt(cadena, 10, 64)
	if err != nil {
		return 0, "Hubo un error "
	}
	return numero, "El número convertido es: " + cadena
}
