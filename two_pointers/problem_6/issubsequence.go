package main

import "fmt"

func isSubSequence(s string, t string) bool {
	if s == "" {
		return true
	}
	i := 0
	for j := 0; j < len(t); j++ {
		if s[i] == t[j] {
			i++
		}
	}
	return i == len(s)
}

func main() {
	s := "abc"
	t := "ahbgdc"
	if isSubSequence(s, t) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
