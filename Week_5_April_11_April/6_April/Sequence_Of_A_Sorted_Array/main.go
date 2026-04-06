package main

import (
	"math"
	"sort"
)

func sortedSquares(nums []int) []int {
	for idx, elem := range nums {
		if elem < 0 {
			elem = -elem
		}
		nums[idx] = elem * elem
	}
	sort.Ints(nums)
	return nums
}

func sortedSquaresTwoPointers(nums []int) []int {
	n := len(nums)
	result := make([]int, n)
	left, right := 0, len(nums)-1
	for i := len(nums) - 1; i >= 0; i-- {
		leftNum := nums[left]
		rightNum := nums[right]

		if math.Abs(float64(leftNum)) > math.Abs(float64(rightNum)) {
			result[i] = leftNum * leftNum
			left++
		} else {
			result[i] = rightNum * rightNum
			right--
		}
	}
	return result
}

func main() {
	nums := []int{-4, -1, 0, 3, 10}
	nums2 := []int{-4, -1, 0, 3, 10}
	result := sortedSquares(nums)
	result2 := sortedSquaresTwoPointers(nums2)
	for _, val := range result {
		println(val)
	}

	for _, val := range result2 {
		println(val)
	}
}
