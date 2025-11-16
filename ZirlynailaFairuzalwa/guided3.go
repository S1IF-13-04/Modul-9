package main
import "fmt"

func main(){
	var bilangan int
	fmt.Scan(&bilangan)
	hasil := bilangan > 0 && bilangan % 2 == 1
	if bilangan < 0 && bilangan % 2 == 0 {
		hasil = true
	}
	fmt.Print(hasil)
}