package main

import "fmt"

type figuras2D interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}
func (r rectangulo) area() float64 {
	return r.base * r.altura
}
func calcular(f figuras2D) {
	fmt.Println("Area: ", f.area())
}

func main() {
	myCuadrado := cuadrado{base: 2}
	myRectangulo := rectangulo{base: 2, altura: 4}

	// Se llama a la funcion calcular la cual utiliza la interfaz figuras2D, de esta forma
	// puede llamar tanto al cuadrado como al rectangulo pese a tener distintos atributos, los dos tienen la misma función area()
	calcular(myCuadrado)
	calcular(myRectangulo)

	// Lista interfaces, nativamente, en go no se puede utilizar arrays de distintos datos. Ya que en go siempre se debe especificar el tipo de dato.
	// Para esto se puede utilizar un interface de slice. De esta forma ya puedo crear mi arreglo con distintos tipos de datos como se muestra
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)
}
