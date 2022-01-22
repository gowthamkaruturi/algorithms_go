package knapsack

func CountSubsets(num []int, sum int) int {
	dp := make([][]int, len(num), sum+1)
	return CountSubSetRecursive(num, dp, sum, 0)
}

func CountSubSetRecursive(num []int, dp [][]int, sum int, currentIndex int) int {

	if sum == 0 {
		return 1
	}

	if len(num) == 0 || currentIndex >= len(num) {
		return 0
	}
	if dp[currentIndex][sum] == 0 {
		sum1 := 0

		if num[currentIndex] <= sum {
			sum1 = CountSubSetRecursive(num, dp, sum-num[currentIndex], currentIndex+1)
		}
		sum2 := CountSubSetRecursive(num, dp, sum, currentIndex+1)
		dp[currentIndex][sum] = sum1 + sum2
	}
	return dp[currentIndex][sum]
}
