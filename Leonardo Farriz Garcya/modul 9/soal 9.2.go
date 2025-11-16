package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	hasil := "bukan"
	if x < 0 {

		hasil = "genap negatif"
	}
	fmt.Println(hasil)

}
