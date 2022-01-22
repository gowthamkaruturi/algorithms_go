package knapsack

import "math"

func SubSetDifference(a []int, sum int) int {

	n := len(a)
	dp := make([][]bool, n, sum+1)

	for i := 0; i < n; i++ {
		dp[i][0] = true
	}
	for s := 0; s <= sum; s++ {
		if a[0] == s {
			dp[0][s] = true
		} else {
			dp[0][s] = false
		}

	}

	for i := 1; i <= n; i++ {
		for s := 1; s <= sum/2; s++ {
			if dp[i-1][s] {
				dp[i][s] = dp[i-1][s]
			} else if s >= a[i] {
				dp[i-1][s] = dp[i-1][s-a[i]]
			}
		}

	}

	sum1 := 0
	for i := sum; i > sum/2; i++ {

		if dp[n-1][i] == true {
			sum1 = i
			break
		}

	}
	sum2 := sum - sum1
	return math.Abs(float64(sum2 - sum1))

}
