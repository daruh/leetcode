package climbing_stairs

func climbStairs(n int) int {

	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	prev_1 := 2
	prev_2 := 1
	var current int

	//dp
	// s1 s2 s3
	var i int
	for i = 2; i < n; i++ {
		current = prev_1 + prev_2
		prev_2 = prev_1
		prev_1 = current
	}

	return current
}

/*	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([]int, n)
	dp[0] = 1
	dp[1] = 2

	//dp
	// s1 s2 s3
	var i int
	for i = 2; i < n; i++ {
		dp[i] = dp[i-2] + dp[i-1]
	}

	return dp[n-1]*/
