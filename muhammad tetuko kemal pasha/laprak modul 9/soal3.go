package main

import "fmt"

func main (){
	var x,y int 
	var a,b int
	fmt.Scan(&x,&y)
	b = x % y
	a = y % x
	if a==0{
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}
	if b==0{
		fmt.Println("true")
	}else {
		fmt.Println("false")
	}
}