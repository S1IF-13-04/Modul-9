package main

import "fmt"

func main() {
    var o, m int
    fmt.Scan(&o)

    m = o / 2
    if o%2 != 0 {
        m = m + 1
    }

    fmt.Println(m)
}
