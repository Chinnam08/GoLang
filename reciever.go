package main

import "fmt"

type Employee struct {
	name string
	age  int
}

func (e *Employee) changeName(newName string) {
	e.name = newName
}

func (e *Employee) changeAge(newAge int) {
	e.age = newAge
}

func main() {
	e := Employee{
		name: "bhargavi",
		age:  22,
	}
	fmt.Printf("Employee name before change %s\n", e.name)
	(&e).changeName("bhagi")
	fmt.Printf("Employee name after change %s\n", e.name)
	fmt.Printf("Employee age before change %d\n", e.age)
	(&e).changeAge(23)
	fmt.Printf("Employee age after change %d\n", e.age)
}
