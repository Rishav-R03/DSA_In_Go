package main

import "fmt"

func largestSumSubArray(arr []int) int {
	maxi := -1_00_00_000
	for i := 0; i < len(arr); i++ { // starting indexes
		prefix := 0
		for j := 0; j < len(arr); j++ {
			prefix += arr[j]
			maxi = max(maxi, prefix)
		}
	}
	return maxi
}

func main() {
	arr := []int{4, -6, 2, 8}
	fmt.Println(largestSumSubArray(arr))
}
