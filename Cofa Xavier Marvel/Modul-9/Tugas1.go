package main

import "fmt"

func main() {
	var bikes, persons int

	fmt.Scan(&persons)
	bikes = (persons / 2) + (persons % 2)
	fmt.Print(bikes)
}
