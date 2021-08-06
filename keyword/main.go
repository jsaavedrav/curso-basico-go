package main

import "fmt"

func main(){
	// Defer: deja la línea en espera y la ejecuta al finalizar toda la funcion.
	// Esto sirve mucho al momento de necesitar cerrar una conexión a una bd.
	defer fmt.Println("Hola")
	fmt.Println("Mundo")

	// Continue y break
	for i := 0; i < 10; i++{
		fmt.Println(i)

		// Continue: La instrucción continue se usa cuando se busca omitir la parte restante del bucle, 
		// volver a la parte superior de este y continuar con una nueva iteración.
		if i == 4 {
			fmt.Println("Es 4")
			continue
		}

		// Break
		if i == 4 {
			fmt.Println("Break")
			break
		}
	}
}