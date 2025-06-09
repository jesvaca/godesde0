package files

import (
	"fmt"
	"os"
)

func GrabaTabla(tabla string) {
	file, err := os.Create("./files/txt/tablas.txt")
	if err != nil {
		fmt.Println("Error al crear el archivo " + err.Error())
		return
		// panic(err)
	} else {
		// file.WriteString(fileName)        pendiente por revisar
		fmt.Fprintln(file, tabla)
		file.Close()
		fmt.Println("Archivo creado " + file.Name() + " correctamente")
	}
}

/*
func SumaTabla()  {
	var fileName string = ejercicios.TablaMultiplicar()

}
*/
