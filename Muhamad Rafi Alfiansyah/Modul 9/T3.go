package main
import "fmt"
func main(){
	var x,y int
	fmt.Scan(&x,&y)
	h := false
	g := false
	if y%x == 0 {
		h = true
	}
	fmt.Println(h)

	if x%y == 0 {
		g = true
	}
	fmt.Print(g)
}