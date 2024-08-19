package main

import "log"

func main() {
	rotationTestCases()
	rearrangingTestCases()
	rangeTestCases()
	subArrayTestCases()
	sortTestCases()
}

// lcd using gcd algorithm.
func lcd(a, b int) int {
	if a == 0 && b == 0 {
		return 0
	}

	return a * b / gcd(a, b)
}

// gcd using Euclides algorithm.
func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

func sortTestCases() {
	log.Println("Sort Cases")
	arr := []int{0, 1, 1, 0, 1, 2, 1, 2, 0, 0, 0, 1}
	result := sortZerosOnesAndTwos(arr)
	log.Println(result)
}

func subArrayTestCases() {
	log.Println("Sub Array Cases")
	arr := []int{-2, -3, 4, -1, -2, 1, 5, -3}
	result := maxSubArray(arr)
	log.Println(result)
}

func rangeTestCases() {
	log.Println("Range Cases")
	arr := []int{5, 7, 5, 2, 10, 12, 11, 17, 14, 1, 44}
	queries := [][]int{{2, 5}, {5, 10}, {0, 10}}
	result := rangeLCMQueries(arr, queries)
	log.Println(result)
}

func rearrangingTestCases() {
	log.Println("Rearranging Cases")

	// Rearranging
	matrix := [][]int{
		{1, 2, 2, 1},
		{1, 3, 2},
		{1, 2, 3, 4, 5, 6, 7},
	}

	for _, arr := range matrix {
		result := rearrangeEvenGreaterThanOdd(arr)
		log.Println(result)
	}

	for _, arr := range matrix {
		result := rearrangeMaximumTwoPointers(arr)
		log.Println(result)
	}
}

func rotationTestCases() {
	log.Println("Rotation Cases")

	// Rotation
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	d := 2

	for i := range d {
		// result := firstApproach(arr, i)
		// result := secondApproach(arr, i)
		result := thirdApproach(arr, i)
		log.Println(i, result)
	}
}
