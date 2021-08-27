package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func Area(r Rectangle) int {
	return r.length * r.width
}

func (r *Rectangle) Area() int {
	return r.length * r.width
}

func main() {
	r := Rectangle{
		length: 12,
		width:  12,
	}

	fmt.Printf("Area %d\n", Area(r))
	p := &r

	fmt.Printf("Area %d\n", p.Area())
}
