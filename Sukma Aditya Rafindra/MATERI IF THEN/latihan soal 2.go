package main

import "fmt"

func main() {
	var n int
	var p string = "positif"

	fmt.Print("masukan nilai: ")
	fmt.Scan(&n)

	if n <= 0 {
		p = "bukan positif"
	}
	fmt.Println(p)
}
