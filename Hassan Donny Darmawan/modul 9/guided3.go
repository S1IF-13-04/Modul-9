package main
import "fmt"

func main (){
	var a int
	fmt.Scan(&a)
	b:= "false"
	
	if a % 2==0 && a<0{
		b= "true"
	}
	fmt.Print(b)
}