package main 
import "fmt"

func main(){
	var o,m,s int64
	fmt.Print("Masukan jumlah orang: ")
	fmt.Scan(&o)

	m= o / 2
	s= o %2

	if s!=0 {
		m += 1
	}
	fmt.Print(m)

}