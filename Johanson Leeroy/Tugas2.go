package main

import "fmt"

func main() {
	var bil int
	var cek string
	fmt.Scan(&bil)
	cek = "bukan"

	if bil < 0 && bil%2 == 0 {
		cek = "genap negatif"
	}
	fmt.Print(cek)
}
