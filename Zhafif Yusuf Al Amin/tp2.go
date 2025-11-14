package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)
	y := "positif"
	if bilangan < 0 {
		y = "bukan positif"
	}
	fmt.Print(y)
}
