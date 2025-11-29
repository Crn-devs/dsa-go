package main

/**
Given a sorted array and a number, write a function called sortedFrequency
that counts the occurrences of the number in the array


    sortedFrequency([1,1,2,2,2,2,3],2) // 4
    sortedFrequency([1,1,2,2,2,2,3],3) // 1
    sortedFrequency([1,1,2,2,2,2,3],1) // 2
    sortedFrequency([1,1,2,2,2,2,3],4) // -1

**/

func SortedFrequency(arr []int, num int) int {

	first := findFirst(arr, num)
	if first == -1 {
		return -1
	}
	last := findLast(arr, num)
	return last - first + 1

}

func findFirst(arr []int, num int) int {

	left := 0
	right := len(arr) - 1
	result := -1

	for left <= right {

		mid := left + (right-left)/2

		if arr[mid] == num {

			result = mid
			// move the window leftwards to find earlier values that may be present
			right = mid - 1
		} else if arr[mid] < num {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return result
}

func findLast(arr []int, num int) int {

	start := 0
	end := len(arr) - 1
	result := -1
	for start <= end {

		mid := start + (end-start)/2

		if arr[mid] == num {
			result = mid
			// move the window rightwards to find later values that may be present
			start = mid + 1
		} else if arr[mid] < num {
			start = mid + 1
		} else {
			end = mid - 1
		}

	}
	return result
}
