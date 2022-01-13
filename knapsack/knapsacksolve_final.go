package knapsack

import (
	"fmt"
	"math"
)

func Knapcsacksolve(profits []int, weights []int, capacity int) int {
	// basic check
	if len(profits) == 0 || len(weights) == 0 || len(profits) != len(weights) || capacity == 0 {
		return 0
	}
	n := len(profits)
	dp := make([]int, capacity+1)
	// single check if there is only one weight
	for c := 0; c <= capacity; c++ {
		if weights[0] <= c {
			dp[c] = profits[0]
		}
	}

	for i := 1; i < n; i++ {
		for c := capacity; c >= 0; c-- {
			profit2 := 0
			profit1 := 0
			if weights[i] <= c {
				profit1 = profits[i] + dp[c-weights[i]]
			}

			profit2 = dp[c]

			dp[c] = int(math.Max(float64(profit1), float64(profit2)))
		}
	}
	fmt.Println(dp)
	return dp[capacity]
}
