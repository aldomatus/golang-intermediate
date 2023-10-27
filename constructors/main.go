package main

import "fmt"

type Employee struct {
	id       int
	name     string
	vacation bool
}

func newEmployee(id int, name string, vacation bool) *Employee {
	return &Employee{
		id:       id,
		name:     name,
		vacation: vacation,
	}
}

func main() {
	// Way to emulate a constructor 1
	e := Employee{}
	fmt.Printf("%v\n", e)

	// Way to emulate a constructor 2
	e2 := Employee{
		id:       1,
		name:     "Aldo",
		vacation: true,
	}
	fmt.Printf("%v\n", e2)

	// Way to emulate a constructor 3
	e3 := new(Employee)
	e3.id = 2
	e3.name = "Aldo Matus"
	e3.vacation = false
	fmt.Printf("%v\n", *e3)

	// Way to emulate a constructor 4
	e4 := newEmployee(1, "Name", true)
	fmt.Println(*e4)
}
