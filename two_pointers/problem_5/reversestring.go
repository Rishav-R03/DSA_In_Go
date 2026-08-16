package main

import "fmt"

func reverseString(s []rune) []rune {
	l, r := 0, len(s)-1
	for l <= r {
		s[l], s[r] = s[r], s[l]
		l++
		r--
	}
	return s
}

func main() {
	s := []rune{'h', 'e', 'l', 'l', 'o'}
	fmt.Println(s)
	s1 := string(reverseString(s))
	fmt.Println(s1)
}
