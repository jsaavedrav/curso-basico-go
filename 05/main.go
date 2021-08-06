package main

import "fmt"

func normalFunction(message string){
	fmt.Println(message)
}

func tripleArgument(a, b int, c string){
	fmt.Println(a, b, c)
}

func returnValue(a int) int {
	return a * 2
}

func dobleReturn(a int) (c, d int){
	return a, a * 2
}

func main(){
	normalFunction("Hola Mundo")
	tripleArgument(1,2,"Hola")
	value := returnValue(2)
	fmt.Println("Value", value)

	value1, value2 := dobleReturn(2)
	fmt.Println("value1 y value2: ",value1, value2)

	valueDescart, _:= dobleReturn(2)
	fmt.Println("valueDesc: ",valueDescart)
}