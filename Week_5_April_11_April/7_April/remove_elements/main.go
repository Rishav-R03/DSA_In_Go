package main

import "fmt"

func removeElement(nums []int, val int) int {
	left := 0
	for right := range nums {
		if nums[right] != val {
			nums[left] = nums[right]
			left++
		}
	}
	return left
}

func main() {
	var size int
	var val int
	fmt.Scan(&size, &val)

	var mat []int
	mat = make([]int, size)

	for row := 0; row < size; row++ {
		fmt.Scan(&mat[row])
	}

	fmt.Printf("Values not equal to val are %d\n", removeElement(mat, val))
	// for row := range size {
	// 	for col := range size {
	// 		fmt.Scanf("%d", &mat[row][col])
	// 	}
	// }

	// for row := range size {
	// 	for col := range size {
	// 		fmt.Printf("%d", &mat[row][col])
	// 	}
	// }

}
