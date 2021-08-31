package main

import "fmt"

func describe(i interface{}) {
	fmt.Printf("Type= %T,value=%v\n", i, i)
}

func main() {
	s := "hello bhagi"
	describe(s)
	name := struct {
		name string
	}{
		name: "bhagi"}

	describe(name)
}
