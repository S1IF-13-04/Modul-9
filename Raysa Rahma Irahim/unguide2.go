package main

import "fmt"

func main() {
	var bil int
	var hasil string

	fmt.Scan(&bil)

	if bil < 0 && bil%2 == 0 {
		hasil = "genap negatif"
	} else {
		hasil = "bukan"
	}
	fmt.Println(hasil)
}
