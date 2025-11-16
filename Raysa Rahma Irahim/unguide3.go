package main

import "fmt"

func main() {
	var x, y int

	fmt.Scan(&x, &y)

	var hasil1, hasil2 bool

	if y%x == 0 {
		hasil1 = true
	} else {
		hasil1 = false
	}

	if x%y == 0 {
		hasil2 = true
	} else {
		hasil2 = false
	}
	fmt.Println(hasil1)
	fmt.Println(hasil2)
}
