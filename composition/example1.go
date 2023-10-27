package main

import "fmt"

type figure interface {
	area() float64
}

type square struct {
	base float64
}

func (s square) area() float64 {
	return s.base * s.base
}

type rectangle struct {
	base   float64
	height float64
}

func (r rectangle) area() float64 {
	return r.base * r.height
}

type ShapeContainer struct {
	shape figure
}

func (sc *ShapeContainer) calculate() float64 {
	return sc.shape.area()
}

func main() {
	rectangle := rectangle{base: 12.0, height: 10.0}
	square := square{base: 10}

	rectangleContainer := ShapeContainer{shape: rectangle}
	squareContainer := ShapeContainer{shape: square}

	fmt.Printf("Rectangle area: %f\n", rectangleContainer.calculate())
	fmt.Printf("Square area: %f\n", squareContainer.calculate())
}
