package main

import (
	"log"
	"slices"
)

func rearrangeEvenGreaterThanOdd(arr []int) []int {
	log.Println("rearrangeEvenGreaterThanOdd", arr)

	n := len(arr)
	for i := 1; i < n; i++ {
		if i%2 == 1 && arr[i] <= arr[i-1] {
			arr[i], arr[i-1] = arr[i-1], arr[i]
		}
	}
	return arr
}

func rearrangeMaximumTwoPointers(arr []int) []int {
	log.Println("rearrangeMaximumTwoPointers", arr)

	slices.Sort(arr)
	n := len(arr)

	var firstPtr int
	lastPtr := n - 1

	var result []int
	movePtr := true
	for {
		if firstPtr > lastPtr {
			break
		}

		if movePtr {
			result = append(result, arr[lastPtr])
			lastPtr--
			movePtr = false
		} else {
			result = append(result, arr[firstPtr])
			firstPtr++
			movePtr = true
		}
	}

	return result
}
