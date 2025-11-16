package main
import "fmt"
func main(){
	var a int
	fmt.Scan(&a)
	x := false
	if a%2==0&&a<0{
		x = true 
	}
	fmt.Print(x)
}