package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)

	m["Jose"] = 14
	m["Pepe"] = 20

	fmt.Println(m)

	// recorrer map (con este método el orden es aleatorio, si se desea orden, se puede usar el slice)
	for i, v := range m {
		fmt.Println(i, v)
	}

	// Encontrar un valor (En caso de que la llave no coincida con un valor, go retornará un "zerovalue",
	// en este caso representado por un cero al ser declarado de tipo dato int )
	value := m["Jose"]
	fmt.Println(value)

	// Encontrar un valor (para poder controlar el caso que no encontremos un valor con la llave enviada,
	// podemos enviar un segundo valor, en este caso representado por ok)
	value2, ok := m["Josep"]
	if ok {
		fmt.Println(value2, ok)
	} else {
		fmt.Println("No se pudo encontrar el valor")
	}
}
