package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)
	hasil := "Positif"

	if x <= 0 {
		hasil = "bukan positif"
	}

	fmt.Println(hasil)
}
