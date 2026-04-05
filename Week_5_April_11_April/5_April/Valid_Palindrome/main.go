package main

func isPalindrome(s string) bool {
	left, right := 0, len(s)-1
	for left < right {
		if !isAlphaNumeric(s[left]) {
			left++
			continue
		}
		if !isAlphaNumeric(s[right]) {
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

func isAlphaNumeric(c byte) bool {
	return (c >= '0' && c <= '9') ||
		(c >= 'A' && c <= 'Z') ||
		(c >= 'a' && c <= 'z')
}

func toLower(c byte) byte {
	if c >= 'A' && c <= 'Z' {
		return c + ('a' - 'A')
	}
	return c
}

func main() {
	s := "A man, a plan, a canal: Panama"
	isValid := isPalindrome(s)
	println(isValid)
}
