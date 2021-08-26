package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter a number")
	fmt.Scanln(&n)
	if n != 1 && n < 5 && n%2 == 0 {
		fmt.Println(n, " is even")
		fmt.Println("Not Weird")
	} else if n > 6 && n < 20 && n%2 == 0 {
		fmt.Println(n, "is even")
		fmt.Println("weird")
	} else if n > 20 && n%2 == 0 {
		fmt.Println(n, "is even")
		fmt.Println("Not weird")
	} else {
		fmt.Println(n, "is odd")
		fmt.Println("weird")
	}
}
