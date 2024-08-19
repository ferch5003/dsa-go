package main

import "math"

func kadanesAlgorithm(arr []int) int {
	var maxEndingHere int
	maxSoFar := math.MinInt

	for _, value := range arr {
		maxEndingHere += value
		maxSoFar = max(maxSoFar, maxEndingHere)
		maxEndingHere = max(maxEndingHere, 0)
	}

	return maxSoFar
}

// maxSubArray O(N), Time Complexity T(N) = 2N
func maxSubArray(arr []int) int {
	return kadanesAlgorithm(arr)
}
