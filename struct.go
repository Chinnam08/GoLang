package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	e1 := &Employee{
		firstName: "bhagi",
		lastName:  "chinnam",
		age:       22}
	fmt.Println("Employee details are:", (*e1).firstName)
	fmt.Println("Employee details are:", (e1).firstName)
}
