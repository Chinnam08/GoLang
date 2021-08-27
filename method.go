package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

type Rectangle struct {
	length int
	width  int
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func (r Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	r := Rectangle{length: 20,
		width: 15,
	}
	fmt.Printf("Area of rectangle %d\n", r.Area())
	c := Circle{radius: 12}
	fmt.Printf("Area of circle %f\n", c.Area())
}
