package main

import "fmt"

func movingZeros(nums []int) {
	//90300
	//393000

	left := 0
	for right := range nums {
		if nums[right] == 0 {
			continue
		} else {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	movingZeros(nums)
	fmt.Print(nums)
}
