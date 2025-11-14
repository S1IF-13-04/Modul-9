package main
import "fmt"

func main() {
	var x, y int
	fmt.Scan(&x, &y)

	var hasilx, hasily bool

	if y%x == 0 {
		hasilx = true 
	}
	if x%y == 0 {
		hasily = true 
	}

	fmt.Println(hasilx, hasily)
}
