package arrays

func Solution(amounts []int, sum int) map[byte][]int {
	windowSum := 0
	windowStart := 0
	m := make(map[byte][]int)
	validCount := 0
	for index, val := range amounts {
		windowSum = windowSum + val
		if windowSum >= sum {
			if windowSum == sum {
				validCount++
				m[byte(validCount)] = amounts[windowStart : index+1]
			}

			windowStart++
		}
		windowSum -= amounts[windowStart]
	}
	return m
}

func AddArray(numbs ...int) int {
	result := 0
	for _, numb := range numbs {
		result += numb
	}
	return result
}
