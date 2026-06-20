package main

import "fmt"

func validPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isAlphanumeric(s[left]) {
			left++
			continue
		}
		if !isAlphanumeric(s[right]) {
			right--
			continue
		}
		if toLower(s[left]) != toLower(s[right]) {
			return false
		}
		left++
		right--
	}
	return true
}

func isAlphanumeric(c byte) bool {
	return (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func main() {
	s := "A man, a plan, a canal: Panama"
	if validPalindrome(s) {
		fmt.Println("The string is a palindrome")
	} else {
		fmt.Println("The string is not a palindrome")
	}
}
