package main

import (
	"fmt"
	"math"
)

type vertex struct {
	a, b float64
}

func (v *vertex) Point(i float64) {
	v.a = v.a * i
	v.b = v.b * i
}

func (v *vertex) Value() float64 {
	return math.Sqrt(v.a*v.a + v.b*v.b)
}

func main() {
	v := vertex{1, 2}
	fmt.Printf(" %f\n", v.Value())
	v.Point(3)
	fmt.Printf("%f\n", v.Value())
}
