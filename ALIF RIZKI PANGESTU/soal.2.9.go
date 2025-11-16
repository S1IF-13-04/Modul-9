package main

import (
	"fmt"
)

func main() {
	var angka int

	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&angka)

	if angka < 0 && angka%2 == 0 {
		fmt.Println("genap negatif")
	} else {
		fmt.Println("bukan")
	}
}
