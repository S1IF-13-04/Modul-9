package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	motor := n / 2

	if n%2 == 1 {
		motor = motor + 1
	}

	fmt.Println(motor)
}
