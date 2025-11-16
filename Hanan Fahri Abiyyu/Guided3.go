package main

import "fmt"

func main() {
  var x int
  var bilangan bool
  fmt.Scan(&x)
  bilangan = false
  if x < 0{
    bilangan = x%2 == 0
    bilangan = true
  }
  fmt.Print(bilangan)
}
