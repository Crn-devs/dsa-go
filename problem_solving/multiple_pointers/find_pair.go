package main

import "slices"

/**
Given an unsorted array and a number n, find if there exists a pair of elements
in the array whose difference is n. This function should return true if the pair exists
or false if it does not.
**/

func FindPair(arr []int, n int) bool {
	if len(arr) <= 0 || n < 0 {
		return false
	}

	slices.Sort(arr)
	// create pointers

	i, j := 0, 1

	for j < len(arr) {

		diff := arr[j] - arr[i]

		if diff == n {
			return true
		} else if diff < n {
			j++
		} else {
			i++
		}

		if i == j {
			j++
		}
	}
	return false
}
