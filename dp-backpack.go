package main

type Chest struct {
	val []int
	wt  []int
}

func Knapsack(chest *Chest, maxWeight int) (int, []int) {
	n := len(chest.val)
	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, maxWeight+1)
	}
	for i := 1; i <= n; i++ {
		for w := 1; w <= maxWeight; w++ {
			if chest.wt[i-1] > w {
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = max(dp[i-1][w], dp[i-1][w-chest.wt[i-1]]+chest.val[i-1])
			}
		}
	}
	selectedItems := make([]int, 0)
	i, w := n, maxWeight
	for i > 0 && w > 0 {
		if dp[i][w] != dp[i-1][w] {
			selectedItems = append(selectedItems, i-1)
			w -= chest.wt[i-1]
		}
		i--
	}
	return dp[n][maxWeight], selectedItems
}
