package arrays

func CanPartition(a []int) bool {

	n := len(a)
	sum := 0

	for i := 0; i < n; i++ {
		sum = sum + a[i]
	}
	if sum%2 != 0 {
		return false
	}
	sum = sum / 2

	dp := make([][]bool, n, sum+1)
	for i := 0; i < n; i++ {
		dp[i][0] = true
	}

	for s := 1; s <= sum; s++ {

		if a[0] == s {
			dp[0][s] = true
		} else {
			dp[0][s] = false
		}
	}

	for i := 1; i < n; i++ {
		for s := 1; s < sum; s++ {

			if dp[i-1][s] {
				dp[i][s] = dp[i-1][s]
			} else if s >= a[i] {
				dp[i][s] = dp[i-1][s-a[i]]
			}

		}
	}

	return dp[n-1][sum]
}
