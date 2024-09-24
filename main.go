package main

//import "fmt"

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	Cmdflags := NewCmdFlags()
	Cmdflags.Execute(&todos)
	todos.toggle(0)
	storage.Save(todos)
}
