package main

import (
    "fmt"
    "strconv"
)

func main() {

    fib := []int{1, 2}
    for fib[len(fib) - 1] < 4000000 {
        fib = append(fib, calc(fib))
    }

    var sum int
    for _, num := range fib {
        if num % 2 == 0 {
            sum += num
        }
    }

    fmt.Println(strconv.Itoa(sum))

}

func calc(s []int) int {
    return s[len(s) - 1] + s[len(s) - 2]
}
