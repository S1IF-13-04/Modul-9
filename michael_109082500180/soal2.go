package main
import "fmt"

func main (){
	var x,y int64
	fmt.Print("masukan bilangannya: ")
	fmt.Scan(&x)

	y = x % 2

	if y ==0 && x < 0 {
		fmt.Print("genap negatif")
	}else{
		fmt.Print("bukan genap negatif")
	}
}