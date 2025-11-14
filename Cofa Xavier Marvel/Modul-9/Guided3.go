package main

import "fmt"

func main() {
	var number int
	var result bool
	fmt.Scan(&number)
	result = number%2 == 0 && number < 0
	fmt.Println(result)
}
