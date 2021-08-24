package main

import "fmt"

func main() {
	names := []string{"bhagi", "lakshmi", "bhargavi"}
	fmt.Println(names)
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[1] = "chinnam"
	fmt.Println(b, a)
	fmt.Println(names)
}
