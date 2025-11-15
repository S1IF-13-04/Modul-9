package main

import "fmt"

func main() {
	var x int

	fmt.Print("masukan bilangan: ")
	fmt.Scan(&x)

	hasil := "bukan positif"

	if x > 0 {
		hasil = "positif"
	}

	fmt.Println(hasil)
}
