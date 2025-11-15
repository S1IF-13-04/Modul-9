package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	hasil := "bukan"

	if n < 0 && n%2 == 0 {
		hasil = "genap negatif"
	}

	fmt.Println(hasil)
}
