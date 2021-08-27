package main

import (
	"fmt"
	"math"
)

type Circle struct {
	radius float64
}

// type Rectangle struct{
// 	length int
// 	width int
// }

func Area(c Circle) float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	c := Circle{
		radius: 12,
	}
	fmt.Printf("Area of Circle %f\n", Area(c))

}
