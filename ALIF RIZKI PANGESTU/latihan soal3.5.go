package main

import "fmt"

func main() {
 var bilangan int
 var hasil bool
 fmt.Print("Masukan bilangan: ")
 fmt.Scan(&bilangan)

 hasil = bilangan % 2 == 0 && bilangan < 0
 fmt.Println(hasil)
}