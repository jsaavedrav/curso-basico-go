package main

import "fmt"

// para poder declarar una clase o struct, debemos realizarlo afuera de la función main
type car struct {
	brand string
	year  int
}

func main() {
	// una opción para instanciar una clase en go
	myCar := car{brand: "Ford", year: 2020}
	fmt.Println(myCar)

	// otra opción para instanciar una clase en go. En caso de no instanciar algún atributo,
	// ocupa el zerovalue que se mencionaba en la clase de maps
	var otherCar car
	otherCar.brand = "Ferrari"
	fmt.Println(otherCar)
}
