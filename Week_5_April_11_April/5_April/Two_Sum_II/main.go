package main

import "fmt"

func BruteForceTwoSumII(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i + 1, j + 1}
			}
		}
	}
	return []int{-1, -1}
}

func TwoSumII(numbers []int, target int) []int {
	left, right := 0, len(numbers)-1

	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return []int{-1, -1}
}

func main() {
	nums := []int{2, 7, 11, 15}
	target := 22
	result := TwoSumII(nums, target)
	resut2 := BruteForceTwoSumII(nums, target)
	fmt.Print(result, "\n", resut2)
}
