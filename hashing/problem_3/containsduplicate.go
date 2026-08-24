package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func containsDuplicate(arr []int) bool {
	if len(arr) == 0 {
		return true
	}

	mp := make(map[int]int)

	for _, val := range arr {
		mp[val]++
	}

	for _, val := range mp {
		if val > 1 {
			return true
		}
	}
	return false
}

func main() {
	inputFile, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Unable to open file: %v", err)
	}
	defer inputFile.Close()

	outputFile, err := os.Create("../output.txt")
	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)

	readInt := func() int {
		if scanner.Scan() {
			var val int
			if _, err := fmt.Sscan(scanner.Text(), &val); err != nil {
				return -1
			}
			return val
		}
		return -1
	}

	writer := bufio.NewWriter(outputFile)
	writer.Flush()

	testCases := readInt()

	for t := range testCases {
		size := readInt()
		arr := make([]int, size)
		for i := range size {
			arr[i] = readInt()
		}
		if containsDuplicate(arr) {
			line := fmt.Sprintf("Case #%d True", t)
			fmt.Println(line)
			if _, err := writer.WriteString(line); err != nil {
				log.Fatalf("failed to write: %v", err)
			}
		} else {
			line := fmt.Sprintf("Case #%d False", t)
			fmt.Println(line)
			if _, err := writer.WriteString(line); err != nil {
				log.Fatalf("failed to write: %v", err)
			}
		}
	}
}
