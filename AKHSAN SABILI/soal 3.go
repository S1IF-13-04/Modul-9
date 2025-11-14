package main
import "fmt"
func main(){
	var x, y int
	var a, b bool
	fmt.Scan(&x,&y)
	if x <= y {
		a = y % x == 0 
	}
	if x >= y {
		b = x % y == 0 
	}
	fmt.Println(a, b)
}