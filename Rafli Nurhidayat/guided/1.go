package main
import "fmt"
func main() {
	var x int

	fmt.Scan(&x)
	hasil := x

	if x < 0 {
		hasil = x * -1
	}
	fmt.Println(hasil)
}