package main

import "fmt"
func main() {
	var a int
	var kata string
	fmt.Scan(&a)
	kata = "bukan"
		if a < 0 && a%-2 == 0 {
		kata = "genap negatif"
		}	
	fmt.Print(kata)
}
