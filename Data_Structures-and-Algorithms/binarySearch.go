package main

import "fmt"

func binarySearch (arr []int, item int) int {
	low := 0
	high := len(arr)
	mid, guess := 0, 0

	for low <= high {
		mid = low + (high - low) / 2
		guess = arr[mid]

		if guess == item {
			return guess
		}

		if guess > item {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return -1
}

func main () {
	numbers := []int{1, 2, 4, 5, 6, 8, 9}
	result := binarySearch(numbers, 7)
	fmt.Println(result)
}