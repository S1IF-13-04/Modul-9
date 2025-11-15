package main

import "fmt"

func main() {
	var x int
	var cek bool
	fmt.Scan(&x)
	if x < 0 {
		cek = x%2 == 0
	}
	fmt.Print(cek)
}
