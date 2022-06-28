package m

func GetSumOfArrays(elements []int, k int) int {

	sum := 0
	for _, element := range elements {
		sum = sum + element
	}
	avg := sum / len(elements)
	return avg

}
