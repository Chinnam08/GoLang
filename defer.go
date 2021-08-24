package main

import "fmt"

func main() {
	var a, k int

	fmt.Println("enter a number")
	fmt.Scan(&a)
	fmt.Println("Enter the max number ")
	fmt.Scan(&k)
	for i := 0; i <= k; i++ {
		if a == 50 {
			defer fmt.Println(i)
		} else {
			fmt.Println(i)
		}
	}
}
