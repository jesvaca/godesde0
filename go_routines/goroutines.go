package go_routines

import (
	"fmt"
	"strings"
	"time"
)

func MiNombreLento(nombre string, cana1 chan bool) {
	letras := strings.Split(nombre, "")
	for _, letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Println(letra)
	}
	cana1 <- true // Se asigna un valor al canal para disparar el evento
}
