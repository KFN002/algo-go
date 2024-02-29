package main

func MaxExpressionValue(nums []int) int {
	n := len(nums)
	dp := make([][]int, n)
	for i := range dp {
		dp[i] = make([]int, n)
	}

	for p := n - 4; p >= 0; p-- {
		for q := p + 1; q < n-2; q++ {
			for r := q + 1; r < n-1; r++ {
				for s := r + 1; s < n; s++ {
					val := nums[s] - nums[r] + nums[q] - nums[p]
					dp[r][s] = max(dp[r][s], val)
				}
			}
		}
	}

	ans := 0
	for r := 0; r < n-1; r++ {
		for s := r + 1; s < n; s++ {
			ans = max(ans, dp[r][s])
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
