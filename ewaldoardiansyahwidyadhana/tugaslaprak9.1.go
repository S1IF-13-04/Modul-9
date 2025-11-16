package main

import "fmt"

func main() {
	var x int
	fmt.Print("jumlah orang: ")
	fmt.Scan(&x)

	if x%2 == 0 {
		x = x / 2
	} else {
		x = x/2 + 1
	}
	fmt.Print("jumlah motor yang dibutuhkan:", x)
}
