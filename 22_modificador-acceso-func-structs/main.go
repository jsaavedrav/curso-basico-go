package main

import (
	"fmt"
	pk "mypackage/mypackage"
	// una forma de trabajar un poco más simple es utilizar los alias. De esta forma es más simple llamar a pk que mypackage
)

func main() {
	var myCar pk.CarPublic
	myCar.Brand = "Ferrari"
	myCar.Year = 2020
	fmt.Println(myCar)

	// var myOtherCar pk.carPrivate
	// fmt.Println(myOtherCar) // no se puede llamar un bloque privado

	pk.PrintMessage("Hooola")

	// pk.printMessage("Hooola") // no se puede llamar un bloque privado
}
