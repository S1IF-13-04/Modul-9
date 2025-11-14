package main
import "fmt"

func main (){
	var b_1,b_2 int
	fmt.Scan(&b_1,&b_2)
	yx:= "false"
	xy:= "false"

	if b_2%b_1==0 {
		yx="true"
	}
	if b_1%b_2==0 {
		xy="true"
	}
	fmt.Println(yx)
	fmt.Print(xy)

}