package main

import "fmt"

func main() {
	var motor, penumpang int
	fmt.Scan(&penumpang)

	if penumpang > 0 {
		motor = (penumpang + 1) / 2
	}
	fmt.Print("Jumlah motor dibutuhkan: ", motor)
}
