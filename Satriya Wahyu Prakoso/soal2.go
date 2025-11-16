package main
import "fmt"
func main() {
	var bil int
	fmt.Scan(&bil)
	if bil < 0 && bil%2 == 0 {
		fmt.Println("genap negatif")
	} else {
		fmt.Println("bukan")
	}
}