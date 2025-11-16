package main
import "fmt"
func main(){
	var x int
	fmt.Scan(&x)
	h := "positif"
	if x<0{
		h = "bukan " + h
	}
	fmt.Print(h)

}