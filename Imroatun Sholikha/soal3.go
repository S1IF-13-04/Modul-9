package main

import "fmt"

func main() {
	var x, y int
	var pertama, kedua bool
	fmt.Scan(&x, &y)

	if y%x == 0 {
		pertama = true
	}

	if x%y == 0 {
		kedua = true
	}

	fmt.Println(pertama, kedua)

}
