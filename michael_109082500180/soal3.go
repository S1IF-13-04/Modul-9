package main
import "fmt"

func main (){
	var x,y int64
	var a,b bool
	fmt.Print("masukan bilangan: ")
	fmt.Scan(&x,&y)

	a= y % x == 0
	b= x % y == 0

	fmt.Println(a)
	fmt.Println(b)
}