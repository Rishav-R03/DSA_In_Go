package main

import (
	"fmt"
	"sort"
)

func FourSumBF(nums []int, target int) [][]int {
	sort.Ints(nums)
	//Bruteforce
	var result [][]int

	n := len(nums)
	if n < 4 {
		return result
	}

	for p := 0; p < n-3; p++ {
		if p > 0 && nums[p] == nums[p-1] {
			continue
		}
		for q := p + 1; q < n-2; q++ {
			if q > p+1 && nums[q] == nums[q-1] {
				continue
			}
			for r := q + 1; r < n-1; r++ {
				if r > q+1 && nums[r] == nums[r-1] {
					continue
				}
				for s := r + 1; s < n; s++ {
					if s > r+1 && nums[s] == nums[s-1] {
						continue
					}
					if nums[p]+nums[q]+nums[r]+nums[s] == target {
						result = append(result, []int{nums[p], nums[q], nums[r], nums[s]})
					}
				}
			}
		}
	}
	return result
}

func BetterSolution(nums []int, target int) [][]int {
	// Twopointers
	sort.Ints(nums)
	n := len(nums)
	var result [][]int
	if n < 4 {
		return result
	}

	for p := 0; p < n-3; p++ {
		if p > 0 && nums[p] == nums[p-1] {
			continue
		}
		for q := p + 1; q < n-2; q++ {
			if q > p+1 && nums[q] == nums[q-1] {
				continue
			}
			//replace the inner two loops with two pointers
			left := q + 1
			right := n - 1

			for left < right {
				curSum := nums[p] + nums[q] + nums[left] + nums[right]

				if curSum == target {
					result = append(result, []int{nums[p], nums[q], nums[left], nums[right]})
					for left < right && nums[left] == nums[left+1] {
						left++
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
	}
	return result
}

func main() {
	nums := []int{1, 0, -1, 0, -2, 2}
	target := 0

	// result := FourSumBF(nums, target)
	result := BetterSolution(nums, target)
	for r := range result {
		fmt.Printf("[")
		for c := range result[0] {
			fmt.Printf(" %d ", result[r][c])
		}
		fmt.Printf("] ")
	}
}
