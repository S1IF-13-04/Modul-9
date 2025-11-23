package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	var b string
	b = ("bukan")
	if a < 0 && a%-2 == 0 {
		b = ("genap negatif")
	}
	fmt.Print(b)
}
