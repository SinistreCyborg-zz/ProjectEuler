package main

import (
    "fmt"
    "strconv"
)

func main() {

    var largestPalindrome int
    for p := 999; p > 0; p-- {
        for q := 999; q > 0; q-- {
            product := p * q
            if isPalindrome(product) && product > largestPalindrome {
                largestPalindrome = product
            }
        }
    }

    fmt.Println(strconv.Itoa(largestPalindrome))

}

// Check if a number is a palindrome.
func isPalindrome(num int) bool {
    str := strconv.Itoa(num)
    if str == reverse(str) {
        return true
    } else {
        return false
    }
}

// Credit to @yazu on StackOverflow
// https://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes) - 1; i < j; i, j = i + 1, j - 1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
