package main

type taskList struct{
	tasks []*task
}

func (t *taskList) agregarALista(tl *task){
	t.tasks = append(t.tasks, tl)
}

type task struct{
	nombre string
	descripcion string
	completado bool
}

func (t *task) marcarCompleta(){
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string){
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string){
	t.nombre = nombre
}

func main(){
	t := &task{
		nombre: "Completar mi curso de platzi",
		descripcion: "Completar mi curso de Go de platzi lo antes posible",
	}
	lista := &taskList{}
}