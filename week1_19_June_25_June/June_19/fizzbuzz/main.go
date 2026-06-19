package main

import (
	"fmt"
	"strconv"
)

func fizzbuzz(n int) []string {
	ans := make([]string, 0, n)
	for i := 1; i <= n; i++ {
		if (i%3 == 0) && (i%5 == 0) {
			ans = append(ans, "fizzbuzz")
		} else if i%3 == 0 {
			ans = append(ans, "fizz")
		} else if i%5 == 0 {
			ans = append(ans, "buzz")
		} else {
			ans = append(ans, strconv.Itoa(i))
		}
	}
	return ans
}

func main() {
	n := 5
	ans := fizzbuzz(n)
	fmt.Println(ans)
}
