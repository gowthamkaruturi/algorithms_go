package knapsack

import "math"

func Knapsack(profits []int, weights []int, capacity int) float64 {
	return knapsackrecursive(profits, weights, capacity, 0)
}

func knapsackrecursive(profits []int, weights []int, capacity int, index int) float64 {

	if capacity <= 0 || index >= len(profits) {
		return 0
	}

	var profit1 float64 = 0

	if weights[index] <= capacity {
		profit1 = float64(profits[index]) + knapsackrecursive(profits, weights, capacity-weights[index], index+1)
	}
	profit2 := knapsackrecursive(profits, weights, capacity, index+1)

	return math.Max(profit1, profit2)
}
