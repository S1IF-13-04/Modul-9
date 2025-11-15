package main
import "fmt"
func main() {
	var x int
	var y bool
	fmt.Scan(&x)
	if x < 0{
		y = x % 2 == 0
	}
	fmt.Print(y)
}