package main

import "fmt"

func main() {
	var a int
	var b string

	fmt.Scan(&a)

	b = "bukan"

	if a < 0 && a%-2 == 0 {
		b = "genap negatif"
	}

	fmt.Print(b)

}
