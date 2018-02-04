package main

import (
	"fmt"
)

const PI = 3.1416

type circle struct {
	redius float32
}

func (c circle) area() float32 {
	return PI * (c.redius * c.redius)
}

type trapezium struct {
	baseA  float32
	baseB  float32
	height float32
}

func (t trapezium) area() float32 {
	return ((t.baseA + t.baseB) / 2) * t.height
}

type ellipse struct {
	a float32
	b float32
}

func (e ellipse) area() float32 {
	return PI * e.a * e.b
}

func main() {
	c := circle{redius: 4}
	t := trapezium{8, 10, 4}
	e := ellipse{2, 3}

	fmt.Println("Area of Circle: ", c.area())
	fmt.Println("Area of Trapezium: ", t.area())
	fmt.Println("Area of ellipse: ", e.area())
}
