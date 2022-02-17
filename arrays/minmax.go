package arrays

import (
	"fmt"
)

func indexOfMinvalue(arr []int32, start int, end int) int {
	if end <= 0 || start < 0 {
		return 0
	}
	minIndex := start
	for i := start + 1; i <= end; i++ {
		if arr[minIndex] > arr[i] {
			minIndex = i
		}
	}

	return minIndex

}
func SelectionSort(arr []int32) {
	var index_minvalue int
	arr_new := arr
	for i := 0; i < len(arr_new); i++ {
		index_minvalue = indexOfMinvalue(arr, i, len(arr)-1)

		temp := arr_new[i]
		arr_new[i] = arr_new[index_minvalue]
		arr_new[index_minvalue] = temp
	}

}
func MinMax(arr []int32) {

	SelectionSort(arr)
	var min_sum, max_sum int64 = 0, 0
	n := len(arr)
	for i, j := 0, n-1; i < n-1; i, j = i+1, j-1 {
		max_sum += int64(arr[j])
		min_sum += int64(arr[i])
	}

	fmt.Println(min_sum, max_sum)
}

// func main() {
// 	arr_int := []int32{256741038, 623958417, 467905213, 714532089, 938071625}
// 	arrays.MinMax(arr_int)
// 	arr_int2 := []int32{256741038, 623958417, 467905213, 714532089, 938071625}

// 	arrays.MinMax(arr_int2)
// }
