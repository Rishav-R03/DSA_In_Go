package main

import "fmt"

func partiotionLables(s string) []int {
	//1. build index maps
	hashMap := make(map[rune]int)
	for ind, val := range s {
		hashMap[val] = ind
	} //TC O(n) SC O(26)

	var result []int
	size, end := 0, 0
	for i, char := range s {
		size++
		if hashMap[char] > end {
			end = hashMap[char]
		}

		if i == end {
			result = append(result, size)
			size = 0
		}
	} // TC O(n) SC O(26)

	//net TC O(n) SC O(26)
	return result
}

func main() {
	fmt.Println("Enter a string")
	var inputString string
	fmt.Scanf("%s", &inputString)

	result := partiotionLables(inputString)

	for _, r := range result {
		fmt.Printf("%d ", r)
	}
}
