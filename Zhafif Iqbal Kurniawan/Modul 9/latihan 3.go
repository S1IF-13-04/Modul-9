package main

import "fmt"

func main() {
	var x, y int
	var a, b string

	fmt.Scan(&x, &y)
	a = "false"
	if y%x == 0 {
		a = "true"
	}

	fmt.Println(a)

	b = "false"
	if x%y == 0 {
		b = "true"
	}

	fmt.Println(b)
}
