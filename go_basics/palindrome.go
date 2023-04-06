package main

import (
    "fmt"
    "strings"
)

func isPalindrome(s string) bool {
    // Convert the string to lowercase and remove spaces
    s = strings.ToLower(strings.ReplaceAll(s, " ", ""))

    // for loop
    for i := 0; i < len(s)/2; i++ {
		// if condition
        if s[i] != s[len(s)-i-1] {
            return false
        }
    }
    return true
}

func main() {
    str := "A man a plan a canal Panama"
    if isPalindrome(str) {
        fmt.Printf("%q is a palindrome\n", str)
    } else {
        fmt.Printf("%q is not a palindrome\n", str)
    }
}
