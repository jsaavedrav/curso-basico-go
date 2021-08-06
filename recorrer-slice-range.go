package main

import (
	"fmt"
	"strings"
)

func isPalindromo(texto string) {
	var textoReversa string
	for i := len(texto) - 1; i >= 0; i-- {
		textoReversa += string(texto[i])
	}
	if strings.ToLower(textoReversa) == strings.ToLower(texto) {
		fmt.Println("Es Palindromo")
	} else {
		fmt.Println("No es Palindromo")
	}
}

func main() {
	slice := []string{"hola", "que", "hace"}
	for i, valor := range slice {
		fmt.Println(i, valor)
	}

	// si no queremos mostrar el indice podemos hacer lo siguiente:
	for _, valor := range slice {
		fmt.Println(valor)
	}

	// si no queremos mostrar el valor podemos hacer lo siguiente:
	for i := range slice {
		fmt.Println(i)
	}

	// Palindromo
	// ama
	// amor a roma
	isPalindromo("Ana")
}
