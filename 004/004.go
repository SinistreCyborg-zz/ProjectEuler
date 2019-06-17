package main

import (
    "strconv"
)

func main() {



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
