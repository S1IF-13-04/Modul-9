package main

import "fmt"

func main() {
	var x, y int
	var cek bool
	fmt.Scan(&x, &y)
	cek = false
	if y%x == 0 {
		cek = true
	}
	fmt.Println(cek)
	cek = false
	if x%y == 0 {
		cek = true
	}
	fmt.Print(cek)

}
