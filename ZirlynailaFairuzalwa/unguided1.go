package main
import "fmt"

func main(){
	var orang int
	fmt.Scan(&orang)
	motor := orang / 2
	sisa := orang % 2
	if sisa != 0 {
		motor++
	}
	fmt.Print(motor)
}