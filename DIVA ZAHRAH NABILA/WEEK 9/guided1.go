package main

import "fmt"

func main() {
	var a float64
	fmt.Scan(&a)
	if a < 0 {
		a = -a
	}
	fmt.Println(a)
}
