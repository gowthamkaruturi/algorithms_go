package arrays

import (
	"math"
)

func LongestUnique(longest string) float64 {

	seen := make(map[string]float64)

	chars := []rune(longest)

	var start, maximumlength float64 = 0, 0

	for end := 0; end < len(chars); end++ {

		if val, ok := seen[string(chars[end])]; ok {

			start = math.Max(start, val+1)

		}
		seen[string(chars[end])] = float64(end)
		maximumlength = math.Max(maximumlength, float64(end)-start+1)

	}

	return maximumlength

}
