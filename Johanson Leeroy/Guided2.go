package main

import "fmt"

func main() {
	var x int
	var hasil string
	hasil = "Bukan Positif"
	fmt.Scan(&x)

	if x > 0 {
		hasil = "Positif"	
	}
	fmt.Print(hasil)
}
