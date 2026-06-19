package main

import "fmt"

func twoSum(arr []int, target int) []int {
	sumMap := make(map[int]int)

	for i, val := range arr {
		diff := target - val
		if j, ok := sumMap[diff]; ok {
			return []int{j, i}
		}
		sumMap[val] = i
	}
	return []int{}
}

func main() {
	arr := []int{3, 5, 6, 9}
	ans := twoSum(arr, 14)
	fmt.Printf("answer is: %d, %d \n", ans[0], ans[1])
}
