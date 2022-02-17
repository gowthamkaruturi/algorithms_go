package arrays

import "fmt"

func PlusMinus(arr []int32) {
	// Write your code here
	totalLength := len(arr)
	var negative, nonNegative, zero int = 0, 0, 0

	for _, val := range arr {
		if val > 0 {
			nonNegative += 1
		} else if val < 0 {
			negative += 1
		} else if val == 0 {
			zero += 1
		}

	}
	var x float64 = float64(nonNegative) / float64(totalLength)
	fmt.Printf("%.6f \n", x)
	fmt.Println(float32(negative / totalLength))
	fmt.Println(float32(zero / totalLength))

}
