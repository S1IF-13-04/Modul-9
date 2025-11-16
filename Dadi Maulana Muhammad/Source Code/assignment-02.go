package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	var hasil string
	hasil = "bukan"
	if x < 0 {
		if x%2 == 0 {
			hasil = "genap negatif"
		}
	}

	fmt.Println(hasil)
}
