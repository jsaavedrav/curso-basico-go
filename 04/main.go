package main

import "fmt"

func main(){
	helloMessage := "Hello"
	worldMessage := "World"

	// Println
	fmt.Println(helloMessage,worldMessage)
	fmt.Println(helloMessage,worldMessage)
	// Printf
	nombre := "Platzi"
	cursos := 500
	fmt.Printf("%s tiene más de %d cursos\n",nombre, cursos)
	fmt.Printf("%v tiene más de %v cursos\n",nombre, cursos) // se puede agregar v en caso de no conocer el tipo de dato.

	// Sprintf
	message := fmt.Sprintf("%s tiene más de %d cursos",nombre, cursos)
	fmt.Println(message)

	// Tipo de datos
	fmt.Printf("helloMessage: %T\n",helloMessage)
	fmt.Printf("cursos: %T",cursos)
}