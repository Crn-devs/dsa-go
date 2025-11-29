package main

/**
Given an array of 1s and 0s which has all 1s first followed by all 0s, write a function
called countZeroes, which returns the number of zeroes in the array.

    countZeroes([1,1,1,1,0,0]) // 2
    countZeroes([1,0,0,0,0]) // 4
    countZeroes([0,0,0]) // 3
    countZeroes([1,1,1,1]) // 0

Time Complexity - O(log n)
**/

func CountZeros(arr []int) int {

	if len(arr) <= 0 {
		return 0
	}

	start, end := 0, len(arr)-1
	firstZero := len(arr)

	for start <= end {

		mid := start + (end-start)/2

		if arr[mid] == 0 {
			firstZero = mid
			end = mid - 1
		} else {
			start = mid + 1
		}
	}
	return len(arr) - firstZero
}
