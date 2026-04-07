package main

import "fmt"

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	writeInd := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[writeInd] = nums[i]
			writeInd++
		}
	}
	return writeInd
}

func main() {
	var size int
	fmt.Print("Enter size of nums")
	fmt.Scanf("%d", &size)
	mat := make([]int, size)

	for row := range size {
		fmt.Scanf("%d", &mat[row])
	}

	fmt.Printf("Number of unqiue elements left %d", removeDuplicates(mat))
}
