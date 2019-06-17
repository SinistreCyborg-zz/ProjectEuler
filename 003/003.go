package main

import (
    "fmt"
    "strconv"
)

func main() {

    i := 2
    num := 600851475143

    for ; num > i; i++ {
        if num % i == 0 {
            num = num / i;
        }
    }

    fmt.Println(strconv.Itoa(i))

}
