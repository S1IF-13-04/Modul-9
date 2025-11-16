package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	hasil := "false"
	if x%2 == 0 && x < 0 {
		hasil = "true"
	}
	fmt.Println(hasil)
}
