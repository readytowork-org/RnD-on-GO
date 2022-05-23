package main

import (
	"fmt"
	"math"
)

// defining the interfaces for various shapes to find the area
type areaInterface interface {
	area() float64
}

type triangle struct {
	base   float64
	height float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	width  float64
	height float64
}

func (t triangle) area() float64 {
	area := (t.base * t.height) / 2
	return area
}

func (r rectangle) area() float64 {
	area := r.width * r.height
	return area
}

func (c circle) area() float64 {
	area := math.Pi * c.radius * c.radius
	return area
}

func main() {

	firstTriangle := triangle{10, 5}
	firstCircle := circle{5.2}
	firstRectangle := rectangle{2.2, 5}

	result := []areaInterface{firstTriangle, firstCircle, firstRectangle}

	// we have set the interface area() and provided it with two different values
	// now, we need to print those values using any loop

	for _, areaInterfaceValue := range result {
		fmt.Println(areaInterfaceValue.area())
	}

	for _, areaInterfaceValue := range result {
		fmt.Println(areaValue(areaInterfaceValue))
	}
}

func areaValue(ar areaInterface) float64 {
	return ar.area()
}
