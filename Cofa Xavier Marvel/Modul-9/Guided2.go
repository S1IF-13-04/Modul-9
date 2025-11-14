package main

import "fmt"

func main() {
	var number int
	var text string
	fmt.Scan(&number)
	text = "non-positive"
	if number > 0 {
		text = "positive"
	}
	fmt.Println(text)
}
