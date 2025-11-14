package main
import "fmt"

func main (){
	var orang,motor int
	fmt.Print("masuka jumlah orang: ")
	fmt.Scan(&orang)

	if orang%2==0 {
		motor = orang/2
	}
	if orang%2==1 {
		motor= (orang/2)+1
	}
	fmt.Print(motor)
}