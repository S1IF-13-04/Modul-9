package main
import "fmt"
func main() {
	var x int
	
	fmt.Scan(&x)
	hasil := "bukan positif"

	if x > 0 {
		hasil = "positif"
	}

	if x <= 0 {
		}
		fmt.Print(hasil)
}