package main

import "fmt"

func main() {
    var JumlahOrang int
	fmt.Print("Jumlah Orang: ")
    fmt.Scan(&JumlahOrang)

    JumlahMotor := (JumlahOrang + 1) / 2
	fmt.Print("Jumlah Motor: ")

    fmt.Println(JumlahMotor)
}