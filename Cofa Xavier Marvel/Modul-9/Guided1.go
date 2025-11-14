package main

import "fmt"

func main() {
	var number int
	fmt.Scan(&number)
	if number < 0 {
		number = -number
	}
	fmt.Println(number)
}
