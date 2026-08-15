package main

import (
	"fmt"
	"unicode"
)

func validPalindrome(s string) bool {
	var newS []rune

	for _, char := range s {
		char = unicode.ToLower(char)
		if isLetter(char) {
			newS = append(newS, char)
		}
	}

	left := 0
	right := len(newS) - 1

	for left < right {
		if newS[left] != newS[right] {
			return false
		}
		left++
		right--
	}
	return true
}

func isLetter(c rune) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c >= '0' && c <= '9')
}

func main() {
	s := "A man, a plan, a canal: Panama"
	if validPalindrome(s) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
