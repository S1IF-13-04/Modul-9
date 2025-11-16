package main

import "fmt"

func main() {
	var x int

	fmt.Scan(&x)

	hasil := x / 2

	if x%2 != 0 {
		hasil = hasil + 1
	}

	fmt.Println(hasil)
}
