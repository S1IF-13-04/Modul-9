package main
import "fmt"

func main() {
    var a int
    fmt.Scan(&a)

    motor := a / 2  
    if a%2 != 0 {   
        motor++
    }

    fmt.Println(motor)
}
