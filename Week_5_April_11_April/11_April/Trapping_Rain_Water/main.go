package main

import "fmt"

func trap(height []int) int {
	n := len(height)
	left := make([]int, n)
	right := make([]int, n)

	max := -1

	for i := 0; i < n; i++ {
		left[i] = max
		if height[i] > max {
			max = height[i]
		}
	}

	max = -1

	for i := n - 1; i > 0; i-- {
		right[i] = max
		if height[i] > max {
			max = height[i]
		}
	}

	water := 0

	for i := 1; i < n-1; i++ {
		min := left[i]
		if right[i] < min {
			min = right[i]
		}
		if min > height[i] {
			water += min - height[i]
		}
	}

	return water
}

func main() {
	var n int
	n, err := fmt.Scanf("%d", &n)
	if err != nil {
		return
	}
	height := make([]int, 0, n)

	for row := range n {
		fmt.Scan(&height[row])
	}

	result := trap(height)
	fmt.Print(result)
}
