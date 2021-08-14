package main

import "fmt"

func message(text string, c chan string) {
	c <- text
}

func main() {
	c := make(chan string, 2)
	c <- "Mensaje1"
	c <- "Mensaje2"
	fmt.Println(len(c), cap(c))
	// close sirve para cerrar el channel y no permitir que se reciba ningún otro dato. Esto permite mejor rendimiento del código en go
	close(c)
	// c <- "Mensaje3"
	// range es una muy buena forma de poder recorrer los valores de los channels aunque estén cerrados, de hecho se recomienda que estén cerrados
	for message := range c {
		fmt.Println(message)
	}
	// en caso de manejar multiples canales, es mejor utilizar Select
	email1 := make(chan string)
	email2 := make(chan string)
	// a continuación no se sabe que channel va a procesar primero, para esto debemos utilizar select. Para esto es imoortante saber
	// cuantos valores vamos a manejar, cuantos channels vamos a utilizar.Comenzamos por medio de un for
	go message("mensaje1", email1)
	go message("mensaje2", email2)
	for i := 0; i < 2; i++ {
		select {
		case m1 := <-email1:
			fmt.Println("Email recibido de email1", m1)
		case m2 := <-email2:
			fmt.Println("Email recibido de email2", m2)
		}
	}

}
