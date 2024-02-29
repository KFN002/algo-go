package main

import (
	"sort"
)

func SortAndMerge(left, right []int) []int {
	sort.Ints(left)
	sort.Ints(right)
	merged := make([]int, 0, len(left)+len(right))
	leftIndex, rightIndex := 0, 0
	for leftIndex < len(left) && rightIndex < len(right) {
		if left[leftIndex] < right[rightIndex] {
			merged = append(merged, left[leftIndex])
			leftIndex++
		} else {
			merged = append(merged, right[rightIndex])
			rightIndex++
		}
	}
	for leftIndex < len(left) {
		merged = append(merged, left[leftIndex])
		leftIndex++
	}
	for rightIndex < len(right) {
		merged = append(merged, right[rightIndex])
		rightIndex++
	}
	return merged
}
