package main
import "fmt"
func main() {
	var x, h int
	fmt.Scan(&x)
	h = x/2
	if x%2 == 1 {
		h = x/2 + x%2
	}
	fmt.Print(h)
}