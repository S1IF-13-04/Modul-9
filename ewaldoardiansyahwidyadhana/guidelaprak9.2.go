package main

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	if x > 0 {
		fmt.Print("positf")
	} else {
		if x < 0 {
			fmt.Print("bukan positif")
		}
	}
}
