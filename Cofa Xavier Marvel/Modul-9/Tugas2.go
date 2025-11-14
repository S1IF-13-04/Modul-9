package main

import "fmt"

func main() {
	var number int
	result := " not"
	fmt.Scan(&number)

	if number%2 == 0 && number < 0 {
		result = " negative even"
	}

	fmt.Println(result)
}
