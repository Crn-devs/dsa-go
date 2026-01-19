package searching

func BinarySearch(arr []int, target int) int {

	// check empty
	if len(arr) <= 0 {
		return -1
	}

	low, high := 0, len(arr)-1

	for low <= high {
		mid := low + (high-low)/2

		// if element at mid == target return it
		if arr[mid] == target {
			return mid

			// if element at mid is greater than the target then drop the ceiling
		} else if arr[mid] > target {
			high = mid - 1
			// else if the element is less than the element at mid bring up the floor
		} else {
			low = mid + 1
		}
	}
	// not in the array
	return -1
}

// a generic version of binary search: takes in a comparrsion function cmp to compare the elements
// this mirrors the slices.binarysearchfunc method in the slices package
func BinarySearchGeneric[T any](arr []T, x T, cmp func(a, b T) int) int {

	if len(arr) <= 0 {
		return -1
	}

	left, right := 0, len(arr)-1

	for left <= right {
		// subtraction for overflow
		mid := left + (right-left)/2
		c := cmp(arr[mid], x)

		if c == 0 {
			// element found return it
			return mid
		}
		if c < 0 {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
