package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func twoSum(arr []int, target int) []int {
	if len(arr) == 0 {
		return []int{-1, -1}
	}
	mp := make(map[int]int)

	for ind, val := range arr {
		diff := target - val
		if prevInd, ok := mp[diff]; ok {
			return []int{prevInd, ind}
		}
		mp[val] = ind
	}
	return []int{-1, -1}
}
func main() {
	file, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Unable to open input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	readInt := func() int {
		if scanner.Scan() {
			var val int
			fmt.Sscan(scanner.Text(), &val)
			return val
		}
		return 0
	}

	// Use readInt() directly to get the count of test cases
	testCases := readInt()

	for t := 1; t <= testCases; t++ {
		size := readInt()
		arr := make([]int, size)
		for i := 0; i < size; i++ {
			arr[i] = readInt()
		}

		target := readInt()

		result := twoSum(arr, target)
		fmt.Printf("Case #%d: %v\n", t, result)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error scanning file: %v\n", err)
	}
}
