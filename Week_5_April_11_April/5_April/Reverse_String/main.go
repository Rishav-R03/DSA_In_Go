package main

func reverseString(s []byte) {
	for i, j := 0, len(s)-1; i < j; i++ {
		s[i], s[j] = s[j], s[i]
		j--
	}
}

func main() {
	s := []byte("hello")
	reverseString(s)
	println(string(s))
}
