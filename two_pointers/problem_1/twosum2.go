package main

import "fmt"

func twoSum2(numbers []int, tar int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == tar {
			return []int{left, right}
		} else if sum > tar {
			right--
		} else {
			left++
		}
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	tar := 9
	result := twoSum2(nums, tar)
	if result[0] != -1 && result[1] != -1 {
		fmt.Printf("the target sum is achieved by: %d and %d\n", nums[result[0]], nums[result[1]])
	} else {

		fmt.Printf("the target sum is not achieved : %d and %d\n", nums[result[0]], nums[result[1]])
	}
}
