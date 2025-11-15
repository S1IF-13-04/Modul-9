package main

import "fmt"

func main() {
	var x int64

	fmt.Print("masukan bilangan: ")
	fmt.Scan(&x)

	hasil := "false"
	if x < 0 {
		x = x % 2

		if x == 0 {
			hasil = "true"
		}
	}
	fmt.Println(hasil)
}
