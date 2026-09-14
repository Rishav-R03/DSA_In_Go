package main

import "fmt"

func maxWealth(arr [][]int) int {
	maxWealth := 0
	for i := 0; i < len(arr); i++ {
		wealth := 0
		for j := 0; j < len(arr[0]); j++ {
			wealth += arr[i][j]
		}

		maxWealth = max(maxWealth, wealth)
	}
	return maxWealth
}
func main() {
	arr := [][]int{{1, 2, 3}, {1, 2, 3}}
	fmt.Println(maxWealth(arr))
}
