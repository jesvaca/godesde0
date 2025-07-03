package deferpanic

import (
	"fmt"
	"log"
)

func VemosDefer() {
	defer fmt.Println("Este es el penultimo mensaje")
	fmt.Println("Este es un primer mensaje")
	defer fmt.Println("Este es el mensaje final")

	fmt.Println("Este es el segundo mensaje")
}

// Panic permite abortar la ejecucion de un sistema con un mensaje
func EjemploPanic() {
	defer func() {
		reco := recover()
		if reco != nil {
			log.Fatalf("Ocurrió un error que generó Panic \n %v", reco)
		}
	}()
	a := 1
	if a == 1 {
		panic("Se encontró el valor de 1")
	}
}
