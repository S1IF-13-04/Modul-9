package main

import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var fx, fy string
	fx = "false"
	fy = "false"

	if y%x == 0 {
		fx = "true"
	}

	if x%y == 0 {
		fy = "true"
	}

	fmt.Println(fx)
	fmt.Println(fy)
}
