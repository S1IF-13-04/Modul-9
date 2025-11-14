package main

import "fmt"

func main() {
	var bilangan int
	fmt.Scan(&bilangan)

	var teks string
	teks = "bukan"
	if bilangan%2 == 0 && bilangan < 0 {
		teks = "genap negatif"
	}
	fmt.Println(teks)
}