package main

import "fmt"

func main() {
	var x int
	fmt.Print("masukan bilangan: ")
	fmt.Scan(&x)

	if x%2 == 0 && x < 0 {
		fmt.Println("genap negatif")
	} else {
		fmt.Println("bukan")
	}
}
