package knapsack

func subsetsummemoization(a []int, sum int) int {
	n := len(a)
	dp := make([]int, n)
	dp[0] = 1
	for s := 1; s <= sum; s++ {
		if a[0] == s {
			dp[s] = 1
		} else {
			dp[s] = 0
		}

	}

	for i := 1; i <= n; i++ {
		for s := sum; s >= 1; s-- {
			if s >= a[i] {
				dp[s] += dp[s-a[i]]
			}
		}

	}
	return dp[sum]

}
