package main

import "fmt"

type pc struct{
	ram int
	disk int
	brand string
}
// Esta funcion permite realizar una impresión del struct de forma personalizada. Permitiendo hacer algo un poco más elegante.
// En resumen es crear una interfaz para imprimir el struct completo de una forma personalizada
func (myPC pc) String() string {
	return fmt.Sprintf("Tengo %d GB RAM, %d GB Disco y es una %s", myPC.ram, myPC.disk, myPC.brand)
}

func main()  {
	myPC := pc{ ram: 16, brand: "msi", disk: 100}
	fmt.Println(myPC)
}