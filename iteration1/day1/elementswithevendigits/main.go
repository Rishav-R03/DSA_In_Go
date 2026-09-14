package main

import "fmt"

func elementsWithEvenDigits(arr []int) int {
	ans := 0
	for _, num := range arr {
		digits := calculateDigits(num)
		if digits%2 == 0 {
			ans++
		}
	}
	return ans
}

func calculateDigits(num int) int {
	digits := 0
	for num > 0 {
		_ = num % 10
		digits++
		num = num / 10
	}
	return digits
}

func main() {
	num := 598
	arr := []int{23, 455, 654, 34566, 5343}
	nums := elementsWithEvenDigits(arr)
	fmt.Println(nums)
	ans := calculateDigits(num)
	fmt.Println(ans)

}
