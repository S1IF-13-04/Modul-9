package main
import "fmt"

func main(){
	var bilangan int
	fmt.Scan(&bilangan)
	hasil := "bukan positif"
	if bilangan > 0 {
		hasil = "positif"
	}
	fmt.Print(hasil)
}