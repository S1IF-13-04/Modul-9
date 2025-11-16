package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var faktorX, faktorY bool

	if y%x == 0 {
		faktorX = true
	}
	if y%x != 0 {
		faktorX = false
	}

	if x%y == 0 {
		faktorY = true
	}
	if x%y != 0 {
		faktorY = false
	}

	fmt.Println(faktorX)
	fmt.Println(faktorY)
}
