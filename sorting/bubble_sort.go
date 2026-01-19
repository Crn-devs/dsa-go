package sorting

import (
	"cmp"
)

// bubble sort non generic

func BubbleSort(arr []int) []int {

	// loop the array
	for i := 0; i < len(arr); i++ {
		// inner loop for comparrison
		for j := 0; j < len(arr)-1-i; j++ {
			// if arr at j is greater than the adjacent element swap them
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	// return the sorted array
	return arr
}

// we want to step through the collection and compare each element with its adjacent element
// if its greater we want to swap them bubbling the larger element to the end of the collection

func BubbleSortGeneric[S ~[]E, E cmp.Ordered](s S) S {

	// optimise
	// using a bool no swap
	var noSwap bool
	// optimise

	if len(s) < 2 {
		return s
	}

	for i := len(s); i > 1; i-- {
		noSwap = true
		for j := 0; j < i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
				noSwap = false
			}
		}
		if noSwap {
			return s
		}
	}
	return s
}
