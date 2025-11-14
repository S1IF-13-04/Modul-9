package main

import "fmt"

func main(){
	var a int
	var b bool
	fmt.Scan(&a)
	if a < 0 {
		b = a % 2 == 0
	}
	fmt.Println(b)
}