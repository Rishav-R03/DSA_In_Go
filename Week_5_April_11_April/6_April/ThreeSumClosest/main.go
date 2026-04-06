package main

import (
	"fmt"
	"math"
	"sort"
)

func TSCBF(nums []int, target int) int {
	sort.Ints(nums)
	answer := nums[0] + nums[1] + nums[2]
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				curSum := nums[i] + nums[j] + nums[k]
				if math.Abs(float64(target-curSum)) < math.Abs(float64(target-answer)) {
					answer = curSum
				}

				if curSum == target {
					return curSum
				}
			}
		}
	}
	return answer
}

func TSCTP(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	result := nums[0] + nums[1] + nums[2]

	for i := 0; i < n; i++ {
		left, right := i+1, len(nums)-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-result)) {
				result = sum
			}
			if sum == target {
				return sum
			} else if sum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}
func main() {
	nums := []int{-1, 2, 1, -4}
	target := 1
	ans := TSCBF(nums, target)
	fmt.Println(ans)
}
