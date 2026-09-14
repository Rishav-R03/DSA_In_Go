package main

import "fmt"

func twoSum(arr []int, target int) []int {
	if len(arr) == 0 {
		return []int{-1, -1}
	}
	mp := make(map[int]int)
	for ind, val := range arr {
		diff := target - arr[ind]
		if prevInd, ok := mp[diff]; ok {
			return []int{prevInd, ind}
		}
		mp[ind] = val
	}
	return []int{-1, -1}
}

func main() {
	arr := []int{1, 2, 3, 4}
	target := 5
	ans := twoSum(arr, target)
	fmt.Printf("Elements %d and %d sum to target %d \n", ans[0], ans[1], target)
}
