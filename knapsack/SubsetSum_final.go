package knapsack

func SubsetSum2(a []int, sum int) bool {

	n := len(a)
	dp := make([]bool, n)
	dp[0] = true
	for s := 1; s <= sum; s++ {
		if a[0] == s {
			dp[s] = true
		} else {
			dp[s] = false
		}

	}

	for i := 1; i <= n; i++ {
		for s := sum; s >= 1; s-- {
			if !dp[s] && s >= a[i] {
				dp[s] = dp[s-a[i]]
			}
		}

	}
	return dp[sum]

}
