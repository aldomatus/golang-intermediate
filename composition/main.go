package main

import "fmt"

// ------------ Interface -----------
type engine interface {
	move()
}

// ------------ Struct and func related -----------
type carEngine struct {
	speedChange int
}

func (ce *carEngine) move() {
	fmt.Printf("car is moving in the %d speed change\n", ce.speedChange)
}

// ------------ Struct and func related -----------
type boatEngine struct {
	numPersons int
}

func (be *boatEngine) move() {
	fmt.Printf("Boat is moving with %d persons\n", be.numPersons)
}

// ------------ Struct inherits of engine interface -----------
type Vehicle struct {
	engine engine
}

func (v *Vehicle) move() {
	v.engine.move()
}

func main() {
	ce := carEngine{speedChange: 1}
	cb := boatEngine{numPersons: 12}

	car := Vehicle{engine: &ce}
	boat := Vehicle{engine: &cb}

	car.move()
	boat.move()
}
