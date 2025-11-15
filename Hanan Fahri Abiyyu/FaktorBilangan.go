package main

import "fmt"

func main() {
	var x, y int
	var b1, b2 bool
	fmt.Scan(&x, &y)
	b1 = false
	if y%x == 0{
		b1 = true
	}
	fmt.Println(b1)
	b2 = false
	if x%y == 0{
		b2 = true
	}
	fmt.Println(b2)
	
}
