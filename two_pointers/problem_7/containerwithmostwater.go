package main

import (
	"fmt"
)

func containerWithMostWater(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	l, r := 0, len(nums)-1
	maxArea := 0
	for l < r {
		width := (r - l)
		height := min(nums[l], nums[r])
		curArea := height * width
		maxArea = max(maxArea, curArea)
		if nums[l] < nums[r] {
			l++
		} else {
			r--
		}
	}
	return int(maxArea)
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	ans := containerWithMostWater(height)
	fmt.Println(ans)
}
