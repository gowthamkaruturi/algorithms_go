package arrays

import "math"

func Knapsackmm(profits []int, weights []int, capacity int) float64 {

	var dp = make([][]float64, len(profits), capacity+1)
	return knapsackmmrecursive(profits, weights, capacity, 0, dp)
}

func knapsackmmrecursive(profits []int, weights []int, capacity int, index int, dp [][]float64) float64 {

	if capacity <= 0 || index >= len(profits) {
		return 0
	}
	if dp[index][capacity] != 0 {
		return dp[index][capacity]
	}
	var profit1 float64 = 0

	if weights[index] <= capacity {
		profit1 = float64(profits[index]) + knapsackrecursive(profits, weights, capacity-weights[index], index+1)
	}
	profit2 := knapsackrecursive(profits, weights, capacity, index+1)

	dp[index][capacity] = math.Max(profit1, profit2)
	return dp[index][capacity]
}
