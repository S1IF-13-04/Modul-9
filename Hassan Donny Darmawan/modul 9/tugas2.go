package main
import "fmt"

func main (){
	var a int
	fmt.Scan(&a)
	b:= "not"
	
	if a<0{
		b= "negative even"
	}
	fmt.Print(b)
}