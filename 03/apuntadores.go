package main

import "fmt"

func main(){
	x := 25
	fmt.Println(x)
	fmt.Println(&x)
	cambiarValor(x)
	fmt.Println(x)
}

func cambiarValor(a int){
	fmt.Println(&a)
	// al imprimir, es un espacio de memoria distinto al main. Entonces, son variables en espacio de memoria distintas
	a = 36
}