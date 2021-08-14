package main

import "fmt"

// chan: se refiere al canal o channel
// una forma para dejar más eficiente el código de nuestros channels, es indicándole en el parámetro si es un channel de entrada o salida.
// para este ejemplo se grafica un channel de entrada, graficado por <-
func say(text string, c chan<- string) {
	c <- text
}
func main() {
	// el segundo parametro es opcional, la cual indica cuantos go routines puede tomar,
	// pero una forma eficiente y ordenada es pasarle este valor. Ya que si no, toma un valor dinámico
	c := make(chan string, 1)

	fmt.Println("Hello")
	go say("Bye", c)

	fmt.Println(<-c)
}
