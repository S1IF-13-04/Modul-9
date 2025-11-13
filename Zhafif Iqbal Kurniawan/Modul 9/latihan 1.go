package main

import "fmt"

func main() {
	var a, b int

	fmt.Scan(&a)

	b = (a / 2)

	if a%2 != 0 {
		b = (a / 2) + 1
	}

	fmt.Print(b)

}
