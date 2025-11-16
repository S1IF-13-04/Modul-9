package main

import "fmt"

func main() {
  var x int
  var teks string
  teks = "positif"
  fmt.Scan(&x)
  if x <= 0 {
  	teks = "Bukan positif"
  }
  fmt.Print(teks)
}
