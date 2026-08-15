package main

import "fmt"

func moveZeroes(nums []int) {
	slow := 0

	for fast := 1; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			nums[slow], nums[fast] = nums[fast], nums[slow]
			slow++
		}
	}
}

func main() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Println(nums)
	moveZeroes(nums)
	fmt.Println(nums)
}
