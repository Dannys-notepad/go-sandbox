package main

import "fmt"

func findBiggest (arr []int) int {
	n := len(arr)
	biggest := arr[0]
	biggestIdx := 0

	for i := 0; i < n; i++ {
		if arr[i] > biggest {
			biggest = arr[i]
			biggestIdx = i
		}
	}

	return biggestIdx
}

func findSmallest (arr []int) int {
	smallest := arr[0]
	smallestIdx := 0
	n := len(arr)

	for i := 0; i < n; i++ {
		if arr[i] < smallest {
			smallest = arr[i]
			smallestIdx = i
		}
	}

	return smallestIdx
}

func selectionSort (arr []int) []int {
	newArr := []int{}
	smallest := 0
	n := len(arr)

	for i := 0; i < n; i++ {
		smallest = findSmallest(arr)
		newArr = append(newArr, arr[smallest])
		arr = append(arr[:smallest], arr[smallest+1:]...)
	}

	return newArr
}

func main () {

	numbers := []int{20, 1, 4, 2, 5, 42, 9, 11, 32, 50, 13, 19}
	//sortedNUmbers := selectionSort(numbers)
	biggestNumber := findBiggest(numbers)
	//fmt.Println(sortedNUmbers)
	fmt.Println(biggestNumber)

}