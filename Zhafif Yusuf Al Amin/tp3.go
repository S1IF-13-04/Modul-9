package main

import "fmt"

func main() {
	var b int
	var h bool

	fmt.Scan(&b)
	h = b%2 == 0 && b < 0
	fmt.Println(h)
}
