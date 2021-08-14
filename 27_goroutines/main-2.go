package main

import (
	"fmt"
	"sync"
	"time"
)

func say(text string, wg *sync.WaitGroup) {
	// defer señala a la go routine que está ok
	defer wg.Done()
	fmt.Println(text)
}

/**
* Los waitgroup son una buena forma de poder realizar concurrencia. Pero no son aconsejable, ya que,
* son poco mantenibles en el tiempo. Para esto existe una instrucción que es más ordenada.
* Nos referimos los channels. Son canales que se pueden comunicar los go routines.
**/
// Para mejorar la función anterior se utilizan las go routine por medio de un waitgroup
func main() {
	// instanciamos la variable wg, la cual contiene la go routine de tipo waitgroup
	var wg sync.WaitGroup
	fmt.Println("Hello")
	// para agregar la go routine, se realiza por medio de wg.Add y agregando el valor 1 indicando el número de go routine a agregar
	wg.Add(1)
	// aquí llamamos a la función say pasando como parametro (puntero) el waitgroup
	go say("World", &wg)
	// aquí indicamos a la go routine que debe esperar que todo los waitgroup terminen
	wg.Wait()

	// funcion anónima
	go func(text string) {
		fmt.Println(text)
	}("Adios")

	time.Sleep(time.Second * 1)
}
