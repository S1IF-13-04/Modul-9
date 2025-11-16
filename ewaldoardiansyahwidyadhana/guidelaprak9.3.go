package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	hasil := x < 0 && x%2 == 0

	fmt.Println(hasil)
}
