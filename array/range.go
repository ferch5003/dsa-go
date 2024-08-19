package main

func multiLCD(arr []int, r int) []int {
	if len(arr) == 1 {
		return []int{lcd(arr[0], r)}
	}

	return multiLCD(multiLCD(arr[:len(arr)-1], arr[len(arr)-1]), r)
}

// rangeLCMQueries O(Q*N) where Q is the number of queries and N is the quantity of elements from the array arr.
// T(N) = Q*(2N + 3)
func rangeLCMQueries(arr []int, queries [][]int) []int {
	var results []int
	for _, query := range queries {
		results = append(results, multiLCD(arr[query[0]:query[1]], arr[query[1]])...)
	}

	return results
}
