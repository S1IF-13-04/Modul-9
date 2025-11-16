package main

import "fmt"

func main() {
	var x int
	var hasil bool

	hasil = false

	fmt.Scan(&x)

	if x%2 == 0 && x < 0 {
		hasil = true
	}

	fmt.Println(hasil)
}
