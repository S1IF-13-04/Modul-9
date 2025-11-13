package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	b := "positif"
	if a <= 0 {
		b = "bukan positif"
	}
	fmt.Print(b)
}
