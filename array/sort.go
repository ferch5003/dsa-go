package main

import "slices"

func sortZerosOnesAndTwos(arr []int) []int {
	counter := map[int][]int{
		0: {},
		1: {},
		2: {},
	}

	for _, value := range arr {
		counter[value] = append(counter[value], value)
	}

	return slices.Concat(counter[0], counter[1], counter[2])
}
