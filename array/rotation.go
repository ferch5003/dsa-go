package main

import "log"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	d := 2

	for i := range d {
		// result := firstApproach(arr, i)
		// result := secondApproach(arr, i)
		result := thirdApproach(arr, i)
		log.Println(i, result)
	}
}

// firstApproach Time O(N) with Space O(N) by using second array.
func firstApproach(arr []int, d int) []int {
	var newArr []int
	n := len(arr)

	ptr := d % n
	linkPtr := (ptr + 1) % n

	newArr = append(newArr, arr[ptr])
	for linkPtr != ptr {
		newArr = append(newArr, arr[linkPtr])
		linkPtr = (linkPtr + 1) % n
	}

	return newArr
}

// secondApproach Time O(d) with Space O(1).
func secondApproach(arr []int, d int) []int {
	if d == 0 {
		return arr
	}

	n := len(arr)

	ptr := d % n
	linkPtr := (ptr + 1) % n

	for linkPtr != ptr {
		arr[ptr], arr[linkPtr] = arr[linkPtr], arr[ptr]
		linkPtr = (linkPtr + 1) % n
	}

	return arr
}

// gcd using Euclides algorithm.
func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}

// thirdApproach Time O(N) with Space O(1).
func thirdApproach(arr []int, d int) []int {
	n := len(arr)

	if d%n == 0 {
		return arr
	}

	distance := gcd(n, d)
	for i := 0; i < distance; i++ {
		temp := arr[i]
		between := i
		for {
			k := between + d
			if k >= n {
				k = k - n
			}

			if k == i {
				break
			}

			arr[between] = arr[k]
			between = k
		}
		arr[between] = temp
	}

	return arr
}
