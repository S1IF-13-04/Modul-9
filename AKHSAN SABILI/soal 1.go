package main
import "fmt"
func main(){
	var n, motor int
	fmt.Scan(&n)
	motor = n / 2
	if n % 2 != 0  {
		motor = (n / 2) + 1
	}
	fmt.Println("jumlah motor :", motor)
}