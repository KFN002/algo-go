package main

func CutCable(price []int, len int) int {
	maxValue := make([]int, len+1)
	for i := 1; i <= len; i++ {
		maxVal := -1
		for j := 1; j <= i; j++ {
			maxVal = max(maxVal, price[j]+maxValue[i-j])
		}
		maxValue[i] = maxVal
	}
	return maxValue[len]
}
