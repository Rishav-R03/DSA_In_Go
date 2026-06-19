package main

func shiftZeroes(arr []int) {
	ind := 0
	for i := range arr {
		if arr[i] != 0 {
			arr[ind], arr[i] = arr[i], arr[ind]
			ind++
		}
	}
}
