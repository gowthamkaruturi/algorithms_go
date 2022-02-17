package arrays

func FirstNonRepeating(arr []int) int {
	var elementFound = -1
	for i, j := 0, 1; i < len(arr)-1 && j < len(arr)-1; i, j = i+1, j+1 {

		if arr[i] == arr[j] {

			continue
		} else {
			elementFound = arr[i+j-1]
			break
		}

	}

	return elementFound

}
