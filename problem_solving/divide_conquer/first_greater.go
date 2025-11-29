package main

/**
Find the index of the FIRST element greater than the target.

Example:
FirstGreater([]int{1,2,4,4,5,7,9}, 4)   // 4 (arr[4] = 5)
FirstGreater([]int{1,2,4,4,5,7,9}, 6)   // 5 (arr[5] = 7)
FirstGreater([]int{1,2,4,4,5,7,9}, 9)   // -1
FirstGreater([]int{2,3,5}, 1)           // 0

**/

func FirstGreaterThan(arr []int, target int) int {

	// using divide and conquer find the first element greater than the target

	// create left + right pointers / bounds

	left := 0
	right := len(arr) - 1
	result := -1
	// create the loop to run while left < right
	for left <= right {

		mid := left + (right-left)/2

		if arr[mid] > target {
			result = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return result
}
