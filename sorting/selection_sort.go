package sorting

import (
	"cmp"
)

func SelectionSort(arr []int) []int {

	if len(arr) <= 1 {
		return arr
	}

	// SELECT a minimum - update when the current value is less than
	for i := 0; i < len(arr); i++ {
		// create the minimum set it to be the first element in the array
		min := i
		// check: is each element in the collection less than the value at minimum
		// [3,4,4,1]
		// [3,4] false, [3,4] false, [3,1] true
		// update min to be j
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		// in the outer loop (once the rest of the values have been checked)
		// swap the minimum value in place of the element we are checking against
		arr[min], arr[i] = arr[i], arr[min]
	}
	return arr
}

func SelectionSortGeneric[S ~[]E, E cmp.Ordered](s S) S {

	var min int

	for i := 0; i < len(s)-1; i++ {
		// set min to be the first elements index
		min = i
		for j := i + 1; j < len(s); j++ {
			// check the first element the element at index(min) againt the rest of the slice
			// if min > next element swap min
			if s[j] < s[min] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
	}
	return s
}
