package main

import "fmt"

func main() {
	var x, y int
	var hasil bool

	fmt.Scan(&x, &y)

	hasil = false

	if y%x == 0 {
		hasil = true
	}
	fmt.Println(hasil)

	hasil = false

	if x%y == 0 {
		hasil = true
	}

	fmt.Println(hasil)

}
