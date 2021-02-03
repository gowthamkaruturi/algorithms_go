package main

import (
	"fmt"
)

func mergeArrays(arr1 []int, arr2 []int) []int {
	arr1position := 0
	arr2position := 0
	mergePosition := 0
	mergeArray := make([]int, len(arr1)+len(arr2))
	for arr1position < len(arr1) && arr2position < len(arr2) {
		if arr1[arr1position] < arr2[arr2position] {

			mergeArray[mergePosition] = arr1[arr1position]
			mergePosition++
			arr1position++

		} else {
			mergeArray[mergePosition] = arr2[arr2position]
			mergePosition++
			arr2position++
		}

	}

	for arr1position < len(arr1) {
		mergeArray[mergePosition] = arr1[arr1position]
		mergePosition++
		arr1position++
	}
	for arr2position < len(arr2) {
		mergeArray[mergePosition] = arr2[arr2position]
		mergePosition++
		arr2position++
	}
	return mergeArray
}
func main() {
	a := []int{2, 4, 6, 8, 10}
	b := []int{1, 7, 11}
	c := mergeArrays(a, b)
	fmt.Println(c)
}
