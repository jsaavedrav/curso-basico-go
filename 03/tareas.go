package main

import "fmt"

type task struct{
	nombre string
	descripcion string
	completado bool
}

func (t task) marcarCompleta(){
	t.completado = true
}

func (t task) actualizarDescripcion(descripcion string){
	t.descripcion = descripcion
}

func (t task) actualizarNombre(nombre string){
	t.nombre = nombre
}

func main(){
	t := task{
		nombre: "Completar mi curso de platzi",
		descripcion: "Completar mi curso de Go de platzi lo antes posible",
	}
	fmt.Printf("%+v\n", t)
	t.marcarCompleta()
	t.actualizarNombre("Finalizar mi curso de GO")
	t.actualizarDescripcion("Finalizarlo lo antes posible")
	fmt.Printf("%+v\n", t)
	// De esta forma no se cambiaron los valores. Esto es por el tema de las referencias. Esto se verá en detalle en el archivo apuntadores.go
}