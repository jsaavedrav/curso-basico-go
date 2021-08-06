package main

import "fmt"

func main(){
	// Array, son inmutables
	var array [4]int
	array[0] = 1
	array[1] = 2
	// len largo array, cap capacidad máxima del array
	fmt.Println(array, len(array), cap(array))
	
	fmt.Println("")
	// slice, son mutables	
	slice := []int{0,1,2,3,4,5,6}
	fmt.Println(slice, len(slice), cap(slice))

	// Métodos en el slice
	fmt.Println("")
	fmt.Println(slice[0])
	fmt.Println(slice[:3]) // hasta el indice 3
	fmt.Println(slice[2:4]) // imprimer del 2 al 4to elemento (el segundo valor es exclusivo)
	fmt.Println(slice[4:]) // imprimer del 4to elemento en adelante
	
	fmt.Println("")
	// Append
	slice = append(slice, 7)
	fmt.Println(slice)

	// append new slice
	newSlice := []int{8, 8, 10}
	// Al poner el newSlice, se addicona los 3 puntos. Esto indica que los elemntos se agregaran independiente por go para asegurar la integridad
	slice = append(slice, newSlice...)
	fmt.Println(slice)
}