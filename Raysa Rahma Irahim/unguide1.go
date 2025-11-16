package main

import "fmt"

func main() {
	var orang int
	var motor int

	fmt.Scan(&orang)

	if orang%2 == 0 {
		motor = orang / 2
	} else {
		motor = (orang / 2) + 1
	}
	fmt.Println("Jumlah motor yang diperlukan:", motor)
}
