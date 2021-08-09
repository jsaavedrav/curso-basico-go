package mypackage

import "fmt"

// CarPublic Car con acceso público
type CarPublic struct {
	Brand string
	Year  int
}

// carPrivate Car con acceso privado
type carPrivate struct {
	brand string
	year  int
}

// PrintMessage es para imprimir un mensaje
func PrintMessage(texto string) {
	fmt.Println(texto)
}

func printMessage(texto string) {
	fmt.Println(texto)
}
