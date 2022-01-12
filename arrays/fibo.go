package arrays

func Fibo(n int, memo []int) int {

	if n <= 2 {
		return n
	}

	if memo[n] != 0 {
		return memo[n]
	}

	memo[n] = Fibo(n-1, memo) + Fibo(n-2, memo)

	return memo[n]
}
