package files

import (
	"bufio"
	"fmt"
	"os"
)

var fileName string = "./files/txt/tablas.txt"

func GrabaTabla(tabla string) {
	// Se crea el archivo con el metodo Create del paquete os
	file, err := os.Create(fileName)
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

func SumaTabla(tabla string) {
	if !Append(tabla) {
		fmt.Println("Error al concatenar contenido")
	}

}

func Append(texto string) bool {
	fmt.Println(fileName)
	archivo, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error al ABRIR el archivo " + err.Error())
		return false
	}
	_, err = archivo.WriteString(texto)
	if err != nil {
		fmt.Println("Error al GRABAR contenido " + err.Error())
		return false
	}
	archivo.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error al leer el archivo " + err.Error())
		return
	}
	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		registro := scanner.Text()
		fmt.Println("> " + registro)
	}
	archivo.Close()
}
