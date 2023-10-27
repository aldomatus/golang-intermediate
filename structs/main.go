package main

import "fmt"

type Employee struct {
	id   int
	name string
}

func (e *Employee) SetId(id int) {
	e.id = id
}

func (e *Employee) SetName(name string) {
	e.name = name
}

func (e *Employee) GetId() int {
	return e.id
}

func (e *Employee) GetName() string {
	return e.name
}

func main() {
	e := Employee{}

	// struct attributes
	e.id = 1
	e.name = "Aldo"

	fmt.Printf("%v\n", e)

	// struct methods
	e.SetId(5)
	fmt.Printf("%v\n", e)

	e.SetName("Matus")
	fmt.Printf("%v\n", e)

	fmt.Printf("Id using a getter: %d\n", e.GetId())
	fmt.Printf("Name using a getter: %s\n", e.GetName())
}
