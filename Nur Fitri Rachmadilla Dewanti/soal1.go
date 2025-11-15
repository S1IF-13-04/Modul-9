package main
import "fmt"
func main() {
	var orang, motor int
	fmt.Scan(&orang)
	motor = orang / 2

	if orang % 2 == 1{
		motor = motor + 1
	}
	fmt.Print(motor)
}