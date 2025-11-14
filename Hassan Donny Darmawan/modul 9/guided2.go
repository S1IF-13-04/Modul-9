package main
import "fmt"

func main () {
	var a int

	fmt.Scan(&a)
	b:= "Positif"
	if a<=0 {
		b="bukan positif"
	}
	fmt.Print(b)
}