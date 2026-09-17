package main

import "fmt"

func prefix(arr []int) int {
	prefixSum := 0
	for i := 0; i < len(arr); i++ {
		prefixSum += arr[i]
	}
	return prefixSum
}

func suffix(arr []int) int {
	suffixSum := 0
	for i := len(arr) - 1; i >= 0; i-- {
		suffixSum += arr[i]
	}
	return suffixSum
}

func listAllSubArrays(arr []int) [][]int {
	var subArrays [][]int
	//iterate all possible starting indexes
	for i := 0; i < len(arr); i++ {
		//iterate all possible ending indexes
		for j := i; j < len(arr); j++ {
			sub := make([]int, j-i+1)
			copy(sub, arr[i:j+1])
			subArrays = append(subArrays, sub)
		}
	}
	return subArrays
}

func divideArray(arr []int) (bool, [][]int) {
	totalSum := 0
	for i := 0; i < len(arr); i++ {
		totalSum += arr[i]
	}
	curSum := 0
	for i := 0; i < len(arr); i++ {
		curSum += arr[i]
		diff := totalSum - curSum
		if diff == curSum {
			return true, [][]int{{0, i}, {len(arr) - i - 1, len(arr) - 1}}
		}
	}
	return false, [][]int{}
}

func main() {
	arr := []int{1, 2, 3, 0}
	fmt.Println(prefix(arr))
	fmt.Println(suffix(arr))
	resultSubArrays := listAllSubArrays(arr)
	for _, subArr := range resultSubArrays {
		fmt.Printf("%v ", subArr)
	}

	ok, result := divideArray(arr)
	if ok {
		fmt.Printf("%d and %d\n", result[0], result[1])
	} else {
		fmt.Printf("No")
	}
}
