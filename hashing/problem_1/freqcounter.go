package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func countFreq(arr []int) map[int]int {
	mp := make(map[int]int)
	for _, val := range arr {
		mp[val]++
	}
	return mp
}

func main() {
	inputFile, err := os.Open("../input.txt")
	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}
	defer inputFile.Close()
	outputFile, err := os.Create("../output.txt")
	if err != nil {
		log.Fatalf("Unable to create file: %v", err)
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	scanner.Split(bufio.ScanWords)

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

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

	testcases := readInt()
	for t := range testcases {
		arrSize := readInt()
		arr := make([]int, arrSize)
		for i := 0; i < arrSize; i++ {
			arr[i] = readInt()
		}

		result := countFreq(arr)

		line := fmt.Sprintf("Case #%d: %v\n", t, result)
		fmt.Println(line)
		_, err := writer.WriteString(line)
		if err != nil {
			log.Fatalf("Failed writing to output %v", err)
		}
	}
}
