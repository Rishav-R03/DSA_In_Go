package main

import (
	"fmt"
	"sort"
)

func ThreeSumBF(nums []int, target int) [][]int {
	var result [][]int
	sort.Ints(nums)
	for i := 0; i < len(nums)-2; i++ {
		for j := i + 1; j < len(nums)-1; j++ {
			for k := j + 1; k < len(nums); k++ {
				if nums[i]+nums[j]+nums[k] == target {
					triplet := []int{nums[i], nums[j], nums[k]}
					result = append(result, triplet)
				}
			}
		}
	}
	return result
}

func ThreeSumTP(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue // skip same numbers
		}
		target := -nums[i]

		left, right := i+1, len(nums)-1

		for left < right {
			curSum := nums[left] + nums[right]
			if curSum == target {
				result = append(result, []int{nums[i], nums[left], nums[right]})
				for left < right && nums[left] == nums[left+1] {
					left++ //ignore same numbers
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if curSum < target {
				left++
			} else {
				right--
			}
		}
	}
	return result
}

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	result := ThreeSumTP(nums)
	for i := 0; i < len(result); i++ {
		fmt.Print("[")
		for j := 0; j < len(result[0]); j++ {
			fmt.Printf(" %d, ", result[i][j])
		}
		fmt.Print("], ")
	}
}
