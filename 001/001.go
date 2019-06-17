package main

import (
    "fmt"
    "strconv"
)

func main() {

    var sum int
    for i := 0; i < 1000; i++ {
        if i % 3 == 0 || i % 5 == 0 {
            sum += i
        }
    }

    fmt.Println(strconv.Itoa(sum))

}
