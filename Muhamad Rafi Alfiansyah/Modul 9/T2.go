package main
import "fmt"
func main(){
	var a int
	fmt.Scan(&a)
	x := "bukan"
	if a%2==0&&a<0{
		x = "genap negatif"
	}
	fmt.Print(x)
}