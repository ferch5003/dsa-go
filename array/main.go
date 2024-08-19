package main

import "log"

func main() {
	rotationTestCases()
	rearrangingTestCases()
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
