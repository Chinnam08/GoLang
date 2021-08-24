package main

import (
	"fmt"
)

func main() {
	empSalary := map[string]int{
		"Bhargavi": 101,
		"Chinnam":  102,
	}
	fmt.Println("map before deletion", empSalary)
	delete(empSalary, "bhargavi")
	fmt.Println("map after deletion", empSalary)

}
