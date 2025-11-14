package main

import "fmt"

func main() {
	var num1, num2 int
	var bol1, bol2 bool

	fmt.Scan(&num1, &num2)

	bol1 = num2%num1 == 0
	bol2 = num1%num2 == 0

	fmt.Print(bol1, bol2)
}
