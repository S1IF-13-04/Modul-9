package main

import "fmt"

func main() {
	var a int

	fmt.Scan(&a)
	h := "bukan"
	if a < 0 && a%2 == 0 {
		h = "genap negatif"
	}
	fmt.Println(h)
}