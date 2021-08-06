package main

import "fmt"

func main()  {
	slice := []string{"hola", "que", "hace"}
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	// si no queremos mostrar el indice podemos hacer lo siguiente:
	for _, valor := range slice {
		fmt.Println(valor)
	}
}