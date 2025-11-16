package main

import "fmt"

func main(){
	var x int
	fmt.Scan(&x)

	var y int

	if x%2 == 0{
	y = x/2
	}

	if x%2 != 0 {
	y = x/2+1
	}

	fmt.Println(y)

}