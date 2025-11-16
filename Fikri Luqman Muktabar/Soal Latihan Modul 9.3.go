package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Masukkan x dan y: ")
    fmt.Scan(&x, &y)

    fmt.Println("Apakah x faktor dari y? =", y%x == 0)
    fmt.Println("Apakah y faktor dari x? =", x%y == 0)
}