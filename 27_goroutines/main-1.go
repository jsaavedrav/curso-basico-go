package main

import (
	"fmt"
	"time"
)

func say(text string) {
	fmt.Println(text)
}

func main() {
	say("Hello")
	go say("World")
	// En este ejemplo, sólo se imprime Hello. Esto es porque "world" se ejecuta en una go routin y al llegar ahí se detiene la ejecución de main.
	// Una forma de demostrar esto es por medio de un time.Sleep. De esta forma se queda esperando un tiempo para terminar la rutina principal
	time.Sleep(time.Second * 1)

}
