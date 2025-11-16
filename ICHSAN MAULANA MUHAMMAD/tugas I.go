package main

import "fmt"

func main() {
	var orang int
	fmt.Scan(&orang)
	if orang%2 == 0 {
		orang = orang / 2
	} else if orang%2 != 0 {
		orang = orang/2 + 1
	}
	fmt.Println(orang)
}
