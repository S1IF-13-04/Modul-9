package main

import "fmt"

func main() {

	var bikes int
	var num int
	position := " Positive"
	evenNegative := " Not Negative Even"
	evenest := false

	evenest = num%2 == 0
	fmt.Scan(&num)
	if num < 0 {
		num *= -1
		position = " Negative"
	}
	if evenest && position == " Negative" {
		evenNegative = " Negative Even "
	}

	bikes = (num / 2) + (num % 2)

	fmt.Print(num, position, evenNegative, bikes)
}
