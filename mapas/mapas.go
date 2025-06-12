package mapas

import "fmt"

func MostrarMapas() {
	paises := make(map[string]string) // Se pueden indicar la longitud del mapa
	fmt.Println(paises)
	paises["Mexico"] = "D.F."
	paises["Argentina"] = "Buenos Aires"
	fmt.Println(paises)
	fmt.Println(paises["Argentina"])

	// Otra forma de crear un mapa
	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       37,
		"Boca Juniors": 30,
	}
	fmt.Println(campeonato)

	// Agregar un elemento al mapa
	campeonato["River Plate"] = 25
	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("El equipo %s, tiene un puntaje de %d \n", equipo, puntaje)
	}

	// Eliminar un elemento del mapa
	delete(campeonato, "Real Madrid")
	fmt.Println(campeonato)

	// Indagar si una clave y dato existe en mapa
	puntaje, existe := campeonato["Toluca"]
	if existe {
		fmt.Printf("El puntaje capturado es %d, y el equipo existe %t \n", puntaje, existe)
	} else {
		fmt.Printf("El equipo no existe %t \n", existe)
	}
}
